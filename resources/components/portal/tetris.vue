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

          const cvs = document.getElementById("gameborder");
          const ctx = cvs.getContext("2d");
         
          const SQ = 60;
          const ROW = 15;
          const COLUMN = 10;
          const VACANT = "pink";
          const bg  = document.getElementById("background");


          function drawSquare(x,y,color){
            ctx.fillStyle = color;
            ctx.fillRect(x*SQ, y*SQ, SQ, SQ);
            ctx.strokeStyle = "white";
            ctx.strokeRect(x*SQ, y*SQ, SQ, SQ);
          }

          
          
          // ctx.drawImage(bg,0,0,600,900);
          // drawSquare(0,0,VACANT);


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