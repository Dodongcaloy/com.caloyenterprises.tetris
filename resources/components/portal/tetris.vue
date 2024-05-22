<template >

    <div  id="wrapper">

    
      <div id="scoretag">
              SCORE : <div id="score">0</div> 
              HI-SCORE : <div id="hiscore">0</div> 
        </div >

        <canvas id = "gameborder"  width="600px" height="1200px">
       </canvas>

         
        

        <div>
            <button id="restart" type="button" class="btn btn-success btn-lrg" >RESTART</button> 
            <button id="pause" type="button" class="btn btn-success btn-lrg shadow-none">PAUSE</button>
        </div> 

        <h1 id="result"></h1>

        <button id="spin" type="button" class="btn btn-success btn-lg shadow-none">ROTATE</button>
                          
        <br>

       <button id="arrowleft" type="button" class="btn btn-success btn-lg shadow-none">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-arrow-left" viewBox="0 0 16 16">
                    <path fill-rule="evenodd" d="M15 8a.5.5 0 0 0-.5-.5H2.707l3.147-3.146a.5.5 0 1 0-.708-.708l-4 4a.5.5 0 0 0 0 .708l4 4a.5.5 0 0 0 .708-.708L2.707 8.5H14.5A.5.5 0 0 0 15 8"/>
              </svg>
        </button>

        <button id="down" type="button" class="btn btn-success shadow-none btn-circle btn-xl">DROP</button>
                  
        <button id="arrowright" type="button" class="btn btn-success btn-lg shadow-none">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-arrow-right" viewBox="0 0 16 16">
                <path fill-rule="evenodd" stroke-width="2" d="M1 8a.5.5 0 0 1 .5-.5h10.793l-2.147-2.146a.5.5 0 1 1 .708-.708l3.5 3.5a.5.5 0 0 1 0 .708l-3.5 3.5a.5.5 0 1 1-.708-.708L12.293 9H1.5A.5.5 0 0 1 1 8z"/>
              </svg>
        </button>
   
  
       
    </div>
    
</template>






<script>

    define(function () {

    return {
        template: template,

        
        mounted: function(){

          
           

              const cvs = document.getElementById("gameborder");
              const ctx = cvs.getContext("2d");
              const scoreElement = document.getElementById("score");
              const resultELement = document.getElementById("result");
              const pauseButton = document.getElementById("pause");
              const restart = document.getElementById("restart");
              const arrowleft = document.getElementById("arrowleft");
              const down = document.getElementById("down");
              const arrowright = document.getElementById("arrowright");
              const spin = document.getElementById("spin");
              const SQ = 60;
              const COLUMN = COL = 10;
              const ROW = 20;
              const VACANT = "#22242A";

              let submitting = true;
              let gameOver = false;
              let score = 0;
              let speed = 900; // Initial speed (milliseconds)*
              


              function drawSquare(x,y,color){
               
                if(!gameOver){
                  ctx.fillStyle = color;
                  ctx.fillRect(x*SQ, y*SQ, SQ, SQ);
                  ctx.strokeStyle = "#1C1E22";
                  ctx.strokeRect(x*SQ, y*SQ, SQ, SQ);
                }

              }


      
              let board = [];

              for (r=0; r<ROW ; r++){

                board[r] = [];
                for(c =0; c < COL; c++){
                    board[r][c] = VACANT;
                }

              }

              function drawBoard(){

                for(r=0; r<ROW ; r++){
                  for (c=0; c<COL; c++){
                    drawSquare(c,r,board[r][c]);
                  }
                }
              }

              drawBoard();


            

              const Z = [ 
                
                [
                  [1,1,0],
                  [0,1,1],
                  [0,0,0],
                ],

                [
                  [0,0,1],
                  [0,1,1],
                  [0,1,0],
                ],

                [
                  [0,0,0],
                  [1,1,0],
                  [0,1,1],
                ],

                [
                  [0,1,0],
                  [1,1,0],
                  [1,0,0],

                ]

              ];

              const I = [
                [
                  [0,0,1,0],
                  [0,0,1,0],
                  [0,0,1,0],
                  [0,0,1,0],
                ],

                [
                  [0,0,0,0],
                  [0,0,0,0],
                  [1,1,1,1],
                  [0,0,0,0],
                ],

                [
                  [0,1,0,0],
                  [0,1,0,0],
                  [0,1,0,0],
                  [0,1,0,0],
                ],

                [
                  [0,0,0,0],
                  [1,1,1,1],
                  [0,0,0,0],
                  [0,0,0,0],
                ],
              ];

              const J = [
                [
                  [0,1,0],
                  [0,1,0],
                  [1,1,0],
                ],

                [
                  [1,0,0],
                  [1,1,1],
                  [0,0,0],
                ],

                [
                  [0,1,1],
                  [0,1,0],
                  [0,1,0],
                ],

                [
                  [0,0,0],
                  [1,1,1],
                  [0,0,1],
                ],
              ];

              const T = [
                [
                  [0,1,0],
                  [1,1,1],
                  [0,0,0],
                ],

                [
                  [0,1,0],
                  [0,1,1],
                  [0,1,0],
                ],

                [
                  [0,0,0],
                  [1,1,1],
                  [0,1,0],
                ],

                [
                  [0,1,0],
                  [1,1,0],
                  [0,1,0],
                ],
              ];

              const L = [
                [
                  [0,1,0],
                  [0,1,0],
                  [0,1,1],
                ],

                [
                  [0,0,0],
                  [1,1,1],
                  [1,0,0],
                ],

                [
                  [1,1,0],
                  [0,1,0],
                  [0,1,0],
                ],

                [
                  [0,0,1],
                  [1,1,1],
                  [0,0,0],
                ],
                
              ];

              const O = [
                [
                  [0,0,0,0],
                  [0,1,1,0],
                  [0,1,1,0],
                  [0,0,0,0],
                ],
              ];

              const S = [
                [
                  [0,1,1],
                  [1,1,0],
                  [0,0,0],
                ],

                [
                  [0,1,0],
                  [0,1,1],
                  [0,0,1],
                ],

                [
                  [0,0,0],
                  [0,1,1],
                  [1,1,0],
                ],

                [
                  [1,0,0],
                  [1,1,0],
                  [0,1,0],
                ]
              ];



                
            // the pieces and their colors

            const PIECES = [
                [Z,"#EE4266"],
                [I,"#8DECB4"],
                [J,"#FF9800"],
                [T,"#9F70FD"],
                [L,"#387ADF"],
                [O,"#FFDB5C"],
                [S,"yellowgreen"],
            ];

            // generate random pieces

              function randomPiece(){
                  let r = randomN = Math.floor(Math.random() * PIECES.length) // 0 -> 6
                  return new Piece( PIECES[r][0],PIECES[r][1]);
              }

              let p = randomPiece();

              // The Object Piece

              function Piece(tetromino,color){
                  this.tetromino = tetromino;
                  this.color = color;
                  
                  this.tetrominoN = 0; // we start from the first pattern
                  this.activeTetromino = this.tetromino[this.tetrominoN];
                  
                  // we need to control the pieces
                  this.x = 4;
                  this.y = -4;
              }

              // fill function

              Piece.prototype.fill = function(color){
                  for( r = 0; r < this.activeTetromino.length; r++){
                      for(c = 0; c < this.activeTetromino.length; c++){
                          // we draw only occupied squares
                          if( this.activeTetromino[r][c]){
                              drawSquare(this.x + c,this.y + r, color);
                          }
                      }
                  }
              }

              // draw a piece to the board

              Piece.prototype.draw = function(){
                  this.fill(this.color);
              }

              // undraw a piece


              Piece.prototype.unDraw = function(){
                  this.fill(VACANT);

              }


              function drawSilhouette() {
                  // Save the current position of the piece
                  let originalX = p.x;
                  let originalY = p.y;

                  // Move the piece down until it collides with something
                  while (!p.collision(0, 1, p.activeTetromino)) {
                      p.y++;
                  }

                  // Draw the silhouette at the current position
                  p.fill("gray");

                  // Restore the original position of the piece
                  p.x = originalX;
                  p.y = originalY;
              }

             
              Piece.prototype.unDrawSilhouette = function() {
                  let silhouetteY = this.y;
                  while (!this.collision(0, 1, this.activeTetromino)) {
                      this.y++;
                  }
                  // Undraw the bottom part of the silhouette
                  this.fill(VACANT);
                  this.y = silhouetteY;
              };

              // move Down the piece

              Piece.prototype.moveDown = function(){
                  if(!this.collision(0,1,this.activeTetromino)){
                      this.unDraw();
                      this.y++;
                      drawSilhouette();
                      this.draw();
                  }else{
                      // we lock the piece and generate a new one
                      this.lock();
                      p = randomPiece();
                  }
                  
              }

              //drop the piece

              Piece.prototype.dropDown = function(){
                  while (!this.collision(0, 1, this.activeTetromino)) {
                      this.unDraw();
                      this.y++;
                      drawSilhouette();
                      this.draw();
                  }
                  
                  // lock the piece and generate a new one
                  this.lock();
                  p = randomPiece();
              }


              // move Right the piece
              Piece.prototype.moveRight = function(){
                  if(!this.collision(1,0,this.activeTetromino)){
                      this.unDrawSilhouette();
                      this.unDraw();
                      this.x++;
                      this.draw();
                  }
              }

              // move Left the piece
              Piece.prototype.moveLeft = function(){
                  if(!this.collision(-1,0,this.activeTetromino)){
                      this.unDrawSilhouette();
                      this.unDraw();
                      this.x--;
                      this.draw();
                  }
              }

              // rotate the piece
              Piece.prototype.rotate = function(){
                  let nextPattern = this.tetromino[(this.tetrominoN + 1)%this.tetromino.length];
                  let kick = 0;
                  
                  if(this.collision(0,0,nextPattern)){
                      if(this.x > COL/2){
                          // it's the right wall
                          kick = -1; // we need to move the piece to the left
                      }else{
                          // it's the left wall
                          kick = 1; // we need to move the piece to the right
                      }
                  }
                  
                  if(!this.collision(kick,0,nextPattern)){
                      this.unDrawSilhouette();
                      this.unDraw();
                      this.x += kick;
                      this.tetrominoN = (this.tetrominoN + 1)%this.tetromino.length; // (0+1)%4 => 1
                      this.activeTetromino = this.tetromino[this.tetrominoN];
                      this.draw();
                  }
              }

          

              // Function to check and increase speed by 20%
              function increaseSpeed() {
                  if (score % 1000 === 0) {
                      speed *= 0.90;
                  }

              }

              // Update score function (assuming it increments score)
              function updateScore() {
                score += 200; // Increment score
                if (score % 1000 === 0) {
                    increaseSpeed(); 
                }
              }

              function lockScoreUpdate(){
                score += 10;
                if (score % 1000 === 0) {
                    increaseSpeed(); 
                }
              }


              //end the game
            
              function endGame(){
                  gameOver = true;
                  pauseButton.style.display = "none";
                  down.style.display = "none";
                  arrowleft.style.display = "none";
                  spin.style.display = "none";
                  arrowright.style.display = "none";
                  restart.innerHTML = "PLAY AGAIN";
                  resultELement.innerHTML = "GAME OVER!";
                  resultELement.classList.add("popup");
                  resultELement.style.opacity = 1;

                  var formData = {
                      Score: score
                  };

                  if (submitting) {
                    $flare.http.post('<% .Helpers.UrlForRoute "score.save"%>', formData)
                        .then(function(response){
                          console.log(response);
                        })
                        .catch(function(error){
                          console.log(error);
                        });
                        submitting = false;
                  }
              }

              Piece.prototype.lock = function(){
                  for( r = 0; r < this.activeTetromino.length; r++){
                      for(c = 0; c < this.activeTetromino.length; c++){
                          // we skip the vacant squares
                          if( !this.activeTetromino[r][c]){
                              continue;
                          }
                          // pieces to lock on top = game over
                          if(this.y + r < 0){
                              endGame();
                              break;
                          }
                         
                          // we lock the piece
                          board[this.y+r][this.x+c] = this.color;
                          lockScoreUpdate();
                          
                      }
                  }
                  

                  // remove full rows
                  for(r = 0; r < ROW; r++){
                      let isRowFull = true;
                      for( c = 0; c < COL; c++){
                          isRowFull = isRowFull && (board[r][c] != VACANT);
                      }
                      if(isRowFull){
                          // if the row is full
                          // we move down all the rows above it
                          for( y = r; y > 1; y--){
                              for( c = 0; c < COL; c++){
                                  board[y][c] = board[y-1][c];
                              }
                          }
                          // the top row board[0][..] has no row above it
                          for( c = 0; c < COL; c++){
                              board[0][c] = VACANT;
                          }

                          // increment the score
                          updateScore();
                        

                      }
                  }
                  // update the board

                  p = randomPiece();
                          drop();
                  
                  drawBoard();
                 
                  
                  // update the score
                  scoreElement.innerHTML = score;
              }

              // collision fucntion

              Piece.prototype.collision = function(x,y,piece){
                  for( r = 0; r < piece.length; r++){
                      for(c = 0; c < piece.length; c++){
                          // if the square is empty, we skip it
                          if(!piece[r][c]){
                              continue;
                          }
                          // coordinates of the piece after movement
                          let newX = this.x + c + x;
                          let newY = this.y + r + y;
                          
                          // conditions
                          if(newX < 0 || newX >= COL || newY >= ROW){
                              return true;
                          }
                          // skip newY < 0; board[-1] will crush our game
                          if(newY < 0){
                              continue;
                          }
                          // check if there is a locked piece alrady in place
                          if( board[newY][newX] != VACANT){
                              return true;
                          }
                      }
                  }
                  return false;
              }

              

              // CONTROL the piece

              document.addEventListener("keydown",CONTROL);

              function CONTROL(event){
                if (!isPaused) {
                  if(event.keyCode == 37){
                      p.moveLeft();
                      p.moveDown();
                      dropStart = Date.now();
                  }else if(event.keyCode == 38){
                      p.rotate();
                      p.moveDown();
                      dropStart = Date.now();
                  }else if(event.keyCode == 39){
                      p.moveRight();
                      p.moveDown();
                      dropStart = Date.now();
                  }else if(event.keyCode == 40){
                      p.dropDown();
                      dropStart = Date.now();
                  }
                }
              }

              let dropStart = Date.now();
              let isPaused = false; // Flag to track if the game is paused or not

              function drop() {
                  if (!isPaused) { // Check if the game is not paused
                      let now = Date.now();
                      let delta = now - dropStart;
                      if (delta > speed) {
                          p.moveDown();
                          dropStart = Date.now();
                      }
                      if (!gameOver) {
                          requestAnimationFrame(drop);
                      }
                  }
              }

              drop();

            

              pauseButton.addEventListener("click", () => {
                  isPaused = !isPaused; // Toggle the pause state
                  if (isPaused) {
                      // Game is paused, so cancel animation frame
                      cancelAnimationFrame(drop);
                      // Change button text to "Play"
                      pauseButton.textContent = "Press to CONTINUE";
                      pauseButton.style.backgroundColor = "#F5DD61";
                      
                     
                  } else {
                      // Game is resumed, so restart animation frame
                      dropStart = Date.now();
                      drop();
                      // Change button text to "Pause"
                      pauseButton.textContent = "PAUSE";
                      pauseButton.style.backgroundColor = "yellowgreen";
                    
                  }
              });


              function resetGame() {
                window.location.reload();
              }
                  
              restart.addEventListener("click", resetGame);

              arrowleft.addEventListener("click",() =>{
                if(!isPaused){
                  p.moveLeft();
                  p.moveDown();
                  dropStart = Date.now();
                }
              });

              down.addEventListener("click",() =>{
                if(!isPaused){
                  p.dropDown();
                  dropStart = Date.now();
                }
              });

              arrowright.addEventListener("click",() =>{
                if(!isPaused){
                  p.moveRight();
                  p.moveDown();
                  dropStart = Date.now();
                }
              });

              spin.addEventListener("click",() =>{
                if(!isPaused){
                  p.rotate();
                  p.moveDown();
                  dropStart = Date.now();
                }
              });

              

              
          
          }
     }
    
    
    })

</script>






<style>




html, body{
  
  margin: 0;
  padding: 0;
  height: 100vh;
  width: 100%;
  overflow: hidden;
}

#wrapper{
    
    position: relative;
    height: 100vh;
    width: 100%;

}

#scoretag{
    border: none;
    border-radius: 10px;
    padding: 10px;
    font-family: 'arial', bold;
    font-weight: bold;
    color:whitesmoke;
    background-color:tan;
    display: inline-block;
}

#score{

    display: inline-block;
    color:aquamarine;
    font-family:'arial', bold;
    margin-right: 20px;

}



#hiscore{

    display: inline-block;
    color:Maroon;
    font-family: 'arial', bold ;
    font-weight: 2000px ;
}


div{
    margin-top: 10px;
    margin-bottom: 10px;
    line-height: 10px;
    font-size: 25px;
    font-weight: bold;
    text-align: center;
   
}


#gameborder{
  display: block;
  border-radius: 5px;
  border: 2px solid black;
  max-width: 80vw;
  max-height: 60vh;
  margin-bottom: 10px;
}

.controls {
  display: absolute;
  justify-content: center;
  align-items: center;
  margin-bottom: 10px;
}

#pause{
  background-color: yellowgreen;
  border: none;
  border-radius: 10px;
  padding: 10px;
  font-family: 'arial', bold;
  font-weight: bold;
  
}



#restart{
  background-color: yellowgreen;
  border: none;
  border-radius: 10px;
  padding: 10px;
  font-family: 'arial', bold;
  font-weight: bold;

}
button:hover{
  cursor: pointer;
}

canvas{
display: inline-block;
margin: 0 auto;
border: 20 px solid black;


}
        
#result{
  color: tomato;
  font-family: 'arial', bold;
  font-weight: bold;
  opacity: 0;
  transition: opacity 1s ease-in-out;
}

#result.popup {
    animation: popup 2s ease-in-out;
}



#spin{
 background-color: gray;
 font-size: large;
 font-family: 'arial', bold;
 font-weight: bold;
 border-radius: 10px;
 border:gray;
}

#arrowleft{
  background-color: gray;
  border-radius: 10px;
  border:gray;
  padding: 15px 20px 10px 20px;
  font-weight: bold;
  font-family: 'arial', bold;
  font-weight: bold;
}


#arrowright{
  background-color: gray;
  border-radius: 10px;
  border: gray;
  padding: 15px 20px 10px 20px;
  font-weight: bold;
}

#down{
  background-color: gray;
  font-family: 'arial', bold;
  font-weight: bold;
  border-radius: 15px;
  border: gray;
}

.btn-circle.btn-xl {
      padding: 15px 20px 10px 20px;
			border-radius: 15px;
			
}



@media screen and (max-width: 768px) {
  #scoretag, #result {
    font-size: 16px;
  }
}

@media screen and (max-width: 576px) {
  #scoretag, #result {
    font-size: 25px;
}

#result {
    font-size: 30px;
}

canvas#gameborder {
    max-width: 90vw;
    max-height: 55vh;
}

.btn-xl {
    padding: 10px 15px;
}

.btn-circle {
    padding: 15px;
}

}


</style>


 


</styl>