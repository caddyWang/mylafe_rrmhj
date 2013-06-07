    $(function(){

        //顶 uid="2"
        $("ul").find(".up").click(function(){
          var up = $(this);
          var down = $(this).siblings(".down");
          var workId = $(this).siblings(".uid").val();
          optUpOrDownOrAttention(up, down, workId, 1, 1);
          
        });
        
        //踩
        $("ul").find(".down").click(function(){
          var up = $(this).siblings(".up");
          var down = $(this);
          var workId = $(this).siblings(".uid").val();
          optUpOrDownOrAttention(up, down, workId, 2, -1);
        });

        //评论
        $("ul").find(".comment").click(function(){
          var workId = $(this).siblings(".uid").val();
          if($('#comments_'+workId).is(':hidden')) {
            $('#comments_'+workId).fadeIn();
          } else {
            $('#comments_'+workId).fadeOut();
          }
        })

        //无限数据读取
        $('#container').scrollPagination({
          'contentPage': '/static/2.html', 
          'contentData': {}, 
          'scrollTarget': $(window), 
          'heightOffset': 10, 
          'beforeLoad': function(){ 
            $('#loading').fadeIn(); 
          },
          'afterLoad': function(elementsLoaded){ 
             $('#loading').fadeOut();
          }
        });

        //菜单指向图片切换效果
        $("#login").mouseover(function(){
          $(this).attr("src","/static/img/login_over.png");
        });
        $("#login").mouseout(function(){
          $(this).attr("src","/static/img/login.png");
        });

        $("#phone").mouseover(function(){
          $(this).attr("src","/static/img/phone_over.png");
        });
        $("#phone").mouseout(function(){
          $(this).attr("src","/static/img/phone.png");
        });

    });


    //顶或踩的动画, type为1表示顶, -1表示踩
    function playPlus(elm, type){
      var div = null,
        left = elm.offset().left - 80,
        top = elm.offset().top - 100;
      if (type == 1)
        div = $('<div class="digg"><table><tr><td>+1</td></tr></table></div>');
      else
        div = $('<div class="undigg"><table><tr><td>-1</td></tr></table></div>');
      div.css({
        left : left,
        top : top
      });
      setTimeout(function(){
        $("body").append(div);
        div.animate({fontSize:100, opacity:0}, 400, function(){
          div.remove();
        });
      }, 50);
    }

    //项、踩动作与后台交互
    function optUpOrDownOrAttention(up, down, workId, optVal, optView){
      if(optView > 0) { 
        var num = up.find(".num");
        num.text(parseInt(num.text())+optView);
        playPlus(up, optView); 
      }
      else { 
        var num = down.find(".num");
        num.text(parseInt(num.text())+optView);
        playPlus(down, optView); 
      }
    }