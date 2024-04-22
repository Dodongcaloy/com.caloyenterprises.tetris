<template>
    <div  >
         SCORE : <div id="score">0</div> 
        <canvas id = "gameborder"  width="600" height="900">
        <img id = "background" src= '<% .Helpers.AssetPath "icons/mountain.jpg" %>'> 

        </canvas> 

    </div>
    
</template>






<script>

    define(function () {

    return {
        template: template,
        mounted: function(){

              console.log("2345s");

              const cvs = document.getElementById("gameborder");
              const ctx = cvs.getContext("2d");
            
              const SQ = 30;
              const COLUMN = 20;
              const ROW = 30;
              const VACANT = "pink";
              const bg  = document.getElementById("background");


              function drawSquare(x,y,color){
                ctx.fillStyle = color;
                ctx.fillRect(x*SQ, y*SQ, SQ, SQ);
                ctx.strokeStyle = "white";
                ctx.strokeRect(x*SQ, y*SQ, SQ, SQ);
              }

              
              
              ctx.drawImage(bg,0,0,600,900);
              // drawSquare(0,0,"black");


              let board = [];

              for (r=0; r<ROW ; r++){
                board[r] = [];
                for(c =0; c < COLUMN; c++){
                    board[r][c] = VACANT;
                }
              }

              function drawBoard(){
                for(r=0; r<ROW ; r++){
                  for (c=0; c<COLUMN; c++){
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



              // let piece = S[3];
              // const pieceColor =  "black";
              // for(r = 0; r < piece.length; r++){
              //     for(c = 0 ; c < piece.length; c++){
              //       if(piece[r][c]){
              //         drawSquare(c+9,r+3,pieceColor);
              //       }
              //     }
              //   }

                
            // the pieces and their colors

            const PIECES = [
                [Z,"red"],
                [I,"green"],
                [J,"yellow"],
                [T,"white"],
                [L,"skyblue"],
                [O,"black"],
                [S,"orange"],
            ];

            // generate random pieces

            function randomPiece(){
                let r = randomN = Math.floor(Math.random() * PIECES.length) // 0 -> 6
                return new Piece( PIECES[r][0],PIECES[r][1]);
            }

            let p = randomPiece();

            // // The Object Piece

            function Piece(tetromino,color) {
                this.tetromino = tetromino;
                this.color = color;
                
                this.tetrominoN = 0; // we start from the first pattern
                this.activeTetromino = this.tetromino[this.tetrominoN];
                
                // we need to control the pieces
                this.x = 9;
                this.y = 3;
            }

            // // fill function

            Piece.prototype.fill = function(color){
              console.log("fill method");
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
              console.log("draw method")  
              this.fill(this.color);
            }

         

            // Piece.prototype.unDraw = function(){
            //     this.fill(VACANT);
            // }

            function Y(x){
              console.log(x)
            }

            Y("banana");

          
          }
     }
    
    
    })

</script>






<style>

 #gameborder{

    border: 1px  solid white;
    
   
 }



 #score{
            display: inline-block;
            color: red;
        }
        div{
            font-size: 25px;
            font-weight: bold;
            font-family: monospace;
            text-align: center;
        }
        canvas{
            display: block;
            margin: 0 auto;
            
        }


</style>