var cols = ["a", "b", "c", "d", "e", "f", "g", "h"];
var rows = ["8", "7", "6", "5", "4", "3", "2", "1"];

const startSquares = {
  a2: { image: "static/pieces/pawn-w.svg", color: "white", name: "PAWN" },
  b2: { image: "static/pieces/pawn-w.svg", color: "white", name: "PAWN" },
  c2: { image: "static/pieces/pawn-w.svg", color: "white", name: "PAWN" },
  d2: { image: "static/pieces/pawn-w.svg", color: "white", name: "PAWN" },
  e2: { image: "static/pieces/pawn-w.svg", color: "white", name: "PAWN" },
  f2: { image: "static/pieces/pawn-w.svg", color: "white", name: "PAWN" },
  g2: { image: "static/pieces/pawn-w.svg", color: "white", name: "PAWN" },
  h2: { image: "static/pieces/pawn-w.svg", color: "white", name: "PAWN" },
  a1: { image: "static/pieces/rook-w.svg", color: "white", name: "ROOK" },
  b1: { image: "static/pieces/knight-w.svg", color: "white", name: "KNIGHT" },
  c1: { image: "static/pieces/bishop-w.svg", color: "white", name: "BISHOP" },
  d1: { image: "static/pieces/queen-w.svg", color: "white", name: "QUEEN" },
  e1: { image: "static/pieces/king-w.svg", color: "white", name: "KING" },
  f1: { image: "static/pieces/bishop-w.svg", color: "white", name: "BISHOP" },
  g1: { image: "static/pieces/knight-w.svg", color: "white", name: "KNIGHT" },
  h1: { image: "static/pieces/rook-w.svg", color: "white", name: "ROOK" },
  a7: { image: "static/pieces/pawn-b.svg", color: "black", name: "PAWN" },
  b7: { image: "static/pieces/pawn-b.svg", color: "black", name: "PAWN" },
  c7: { image: "static/pieces/pawn-b.svg", color: "black", name: "PAWN" },
  d7: { image: "static/pieces/pawn-b.svg", color: "black", name: "PAWN" },
  e7: { image: "static/pieces/pawn-b.svg", color: "black", name: "PAWN" },
  f7: { image: "static/pieces/pawn-b.svg", color: "black", name: "PAWN" },
  g7: { image: "static/pieces/pawn-b.svg", color: "black", name: "PAWN" },
  h7: { image: "static/pieces/pawn-b.svg", color: "black", name: "PAWN" },
  a8: { image: "static/pieces/rook-b.svg", color: "black", name: "ROOK" },
  b8: { image: "static/pieces/knight-b.svg", color: "black", name: "KNIGHT" },
  c8: { image: "static/pieces/bishop-b.svg", color: "black", name: "BISHOP" },
  d8: { image: "static/pieces/queen-b.svg", color: "black", name: "QUEEN" },
  e8: { image: "static/pieces/king-b.svg", color: "black", name: "KING" },
  f8: { image: "static/pieces/bishop-b.svg", color: "black", name: "BISHOP" },
  g8: { image: "static/pieces/knight-b.svg", color: "black", name: "KNIGHT" },
  h8: { image: "static/pieces/rook-b.svg", color: "black", name: "ROOK" },
};

var userMove = {};
var selectedPiece;

function addSquares() {
  let board = document.getElementById("board");
  for (let row = 0; row < 8; row++) {
    for (let col = 0; col < 8; col++) {
      let sq = document.createElement("div");
      sq.classList.add("square");
      sq.dataset.row = row;
      sq.dataset.col = col;
      if ((row + col) % 2 === 0) {
        sq.classList.add("light");
      } else {
        sq.classList.add("dark");
      }
      sq.id = cols[col] + rows[row];
      if (sq.id in startSquares) {
        let image = `url(${startSquares[sq.id]["image"]})`;
        let color = startSquares[sq.id]["color"];
        let name = startSquares[sq.id]["name"];
        let piece = createPiece(image, color, name);
        sq.appendChild(piece);
      } else {
        let nullSq = document.createElement("div");
        nullSq.classList.add("piece");
        nullSq.id = "NULL";
        nullSq.dataset.color = "none";
        nullSq.style.cursor = "pointer";
        sq.appendChild(nullSq);
      }
      board.appendChild(sq);
    }
  }
}

function createPiece(image, color, name) {
  let piece = document.createElement("div");
  piece.classList.add("piece");
  piece.style.backgroundImage = image;
  piece.dataset.color = color;
  piece.dataset.name = name;
  return piece;
}

function play() {
  resetBoard();
  addSquares();
  let pieces = document.querySelectorAll(".piece");
  pieces.forEach((piece) => {
    piece.addEventListener("click", function () {
      selectSquare(piece);
    });
    piece.style.cursor = "pointer";
  });

  let url = "http://localhost:3435/play";
  fetch(url)
    .then((response) => {
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      return response.json(); // Assuming server returns JSON data
    })
    .then((data) => {
      if (data["color"] == "black") {
        flipBoard();
        updateBoard(data);
      }
      let user = document.createElement("div");
      user.id = "user";
      user.dataset.color = data["color"];
      document.querySelector("body").appendChild(user);
    })
    .catch((error) => {
      console.error("Error fetching data:", error);
    });
}

function selectSquare(piece) {
  const user = document.getElementById("user");
  if (piece.dataset.color === user.dataset.color) {
    setUserMoveFromSq(piece);
  } else {
    if (Object.keys(userMove).length > 0) {
      setUserMoveToSq(piece);
      sendMove();
      userMove = {};
    }
  }
}

function setUserMoveFromSq(piece) {
  let sq = piece.parentNode;
  userMove["from"] = [parseInt(sq.dataset.row), parseInt(sq.dataset.col)];
  selectedPiece = piece;
}

function setUserMoveToSq(piece) {
  let sq = piece.parentNode;
  userMove["to"] = [parseInt(sq.dataset.row), parseInt(sq.dataset.col)];
}

function sendMove() {
  const url = "http://localhost:3435/usermove";

  let move = {
    from: userMove["from"],
    to: userMove["to"],
    promotion: "",
  };

  if (canPromotePawn(userMove)) {
    move["promotion"] = "QUEEN";
  }
  const fetchOptions = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(move), // Convert JavaScript object to JSON string
  };

  fetch(url, fetchOptions)
    .then((response) => {
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      return response.json(); // Parse the JSON data returned by the server
    })
    .then((data) => {
      console.log(data["receipt"]);
      if (data["valid"]) {
        console.log(data["fen"]);
        updateBoard(data);
        getBotMove();
      }
    })
    .catch((error) => {
      console.error("Error fetching data:", error);
    });
}

function canPromotePawn(move) {
  let piece = getFromPiece(move);
  if (piece.dataset.name === "PAWN") {
    let toSq = getToSquare(move);
    if (piece.dataset.color == "white" && parseInt(toSq.dataset.row) == 0) {
      return true;
    }
    if (piece.dataset.color == "black" && parseInt(toSq.dataset.row) == 7) {
      return true;
    }
  }
  return false;
}

function getFromPiece(move) {
  let rank = rows[move["from"][0]];
  let file = cols[move["from"][1]];
  return document.getElementById(file + rank).firstChild;
}

function getToSquare(move) {
  let rank = rows[move["to"][0]];
  let file = cols[move["to"][1]];

  return document.getElementById(file + rank);
}

function getBotMove() {
  const url = "http://localhost:3435/botmove";

  fetch(url)
    .then((response) => {
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      return response.json(); // Assuming server returns JSON data
    })
    .then((data) => {
      console.log(data["receipt"]);
      console.log(data["fen"]);
      updateBoard(data);
    })
    .catch((error) => {
      console.error("Error fetching data:", error);
    });
}

function updateBoard(data) {
  if (data["type"] === "CHECKMATE") {
    let msg = "Checkmate\n" + data["winner"] + " has won!";
    alert(msg);
    location.reload();
  }
  if (data["type"] === "STALEMATE") {
    let msg = "Stalemate. Game ends in a draw";
    alert(msg);
    location.reload();
  }
  let fromRow = data["from"][0];
  let fromCol = data["from"][1];
  let toRow = data["to"][0];
  let toCol = data["to"][1];

  let fromSq = document.getElementById(cols[fromCol] + rows[fromRow]);
  let toSq = document.getElementById(cols[toCol] + rows[toRow]);

  let fromPiece = fromSq.querySelector(".piece");
  let toPiece = toSq.querySelector(".piece");

  let nullPiece = newNullPiece();

  if (data["type"] === "ENPASSANT") {
    handleEnPassant(data);
  }
  if (data["type"] === "CASTLE") {
    handleCastle(data);
  }

  toSq.appendChild(fromPiece);
  toSq.removeChild(toPiece);
  fromSq.appendChild(nullPiece);

  if (data["promotion"] === "QUEEN") {
    handlePawnPromotion(toSq, fromPiece);
  }

  updateEvalBar(data["eval"]);
}

function updateEvalBar(eval) {
  let rounded = Math.round(eval);
  let blackBar = document.getElementById("black-bar");
  let height = (20 - rounded) * 2.5;
  blackBar.style.height = height + "%";
}

function handlePawnPromotion(toSq, fromPiece) {
  let color = fromPiece.dataset.color;
  var image;
  if (color === "white") {
    image = `url("static/pieces/queen-w.svg")`;
  } else {
    image = `url("static/pieces/queen-b.svg")`;
  }
  let queen = createPiece(image, color, "QUEEN");
  queen.addEventListener("click", function () {
    selectSquare(queen);
  });
  queen.style.cursor = "pointer";

  toSq.removeChild(toSq.firstChild);
  toSq.appendChild(queen);
}

function handleEnPassant(data) {
  let epNullPiece = newNullPiece();
  let epCaptureRow = data["from"][0];
  let epCaptureCol = data["to"][1];
  let captureSq = document.getElementById(
    cols[epCaptureCol] + rows[epCaptureRow],
  );
  captureSq.removeChild(captureSq.firstChild);
  captureSq.appendChild(epNullPiece);
}

function handleCastle(data) {
  if (data["color"] === "WHITE") {
    if (cols[data["to"][1]] === "g") {
      let f1 = document.getElementById("f1");
      let rookSq = document.getElementById("h1");
      let rook = rookSq.firstChild;
      rookSq.appendChild(f1.firstChild);
      f1.appendChild(rook);
    } else {
      let d1 = document.getElementById("d1");
      let rookSq = document.getElementById("a1");
      let rook = rookSq.firstChild;
      rookSq.appendChild(d1.firstChild);
      d1.appendChild(rook);
    }
  } else {
    if (cols[data["to"][1]] === "g") {
      let f8 = document.getElementById("f8");
      let rookSq = document.getElementById("h8");
      let rook = rookSq.firstChild;
      rookSq.appendChild(f8.firstChild);
      f8.appendChild(rook);
    } else {
      let d8 = document.getElementById("d8");
      let rookSq = document.getElementById("a8");
      let rook = rookSq.firstChild;
      rookSq.appendChild(d8.firstChild);
      d8.appendChild(rook);
    }
  }
}

function newNullPiece() {
  let nullPiece = document.createElement("div");
  nullPiece.classList.add("piece");
  nullPiece.id = "NULL";
  nullPiece.addEventListener("click", function () {
    selectSquare(nullPiece);
  });
  nullPiece.dataset.color = "none";
  nullPiece.style.cursor = "pointer";
  return nullPiece;
}

function flipBoard() {
  board = document.getElementById("board");
  const squares = Array.from(board.children);
  squares.reverse().forEach((square) => board.appendChild(square));
}

function resetBoard() {
  let board = document.getElementById("board");
  while (board.firstChild) {
    board.removeChild(board.firstChild);
  }
}

addSquares();
