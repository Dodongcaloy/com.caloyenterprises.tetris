<template>
    <div id="background" >
       Score : <div id="score">0</div> 
      <canvas id="gameborder"  width="300" height="600"> </canvas> 
      <!-- <canvas id="myCanvas" width="300" height="150" style="border:1px solid grey"></canvas> -->

    </div>
    
</template>






<script>

    define(function () {

    return {
        template: template,
        mounted: function(){

          const canvas = document.getElementById("gameborder");
          const ctx = canvas.getContext("2d");
         
          const SQ = 30;
          const ROW = 20;
          const COLUMN = 10;
          const VACANT = "white";
  


          function drawSquare(x,y,color){
            ctx.fillStyle = color;
            ctx.fillRect(x*SQ, y*SQ, SQ, SQ);
            ctx.strokeStyle = "black";
            ctx.strokeRect(x*SQ, y*SQ, SQ, SQ);
          }


          // let board = [];

          // for (r=0; r<ROW ; r++){
          //   board[r] = [];
          //   for(c =0; c < COLUMN; c++){
          //       board[r][c] = VACANT;
          //   }
          // }

          // function drawBoard(){
          //   for(r=0; r<ROW ; r++){
          //     for (r=0; r<COLUMN; r++){
          //       drawSquare(c,r,board[r][c]);
          //     }
          //   }
          // }
          
          // drawBoard()

 
        }
     }
    
    
    })

</script>






<style>

 #gameborder{

    border: 1px  solid black;
    
   
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