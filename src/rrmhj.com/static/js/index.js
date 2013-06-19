    Date.prototype.format = function(format)
    {
        var o = {
        "M+" : this.getMonth()+1, //month
        "d+" : this.getDate(),    //day
        "h+" : this.getHours(),   //hour
        "m+" : this.getMinutes(), //minute
        "s+" : this.getSeconds(), //second
        "q+" : Math.floor((this.getMonth()+3)/3),  //quarter
        "S" : this.getMilliseconds() //millisecond
        }
        if(/(y+)/.test(format)) format=format.replace(RegExp.$1,
        (this.getFullYear()+"").substr(4 - RegExp.$1.length));
        for(var k in o)if(new RegExp("("+ k +")").test(format))
        format = format.replace(RegExp.$1,
        RegExp.$1.length==1 ? o[k] :
        ("00"+ o[k]).substr((""+ o[k]).length));
        return format;
    }


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

        //展现评论
        $("ul").find(".comment").click(function(){
          var workId = $(this).siblings(".uid").val();
          if($('#comments_'+workId).is(':hidden')) {
            var hasComment = $('#commlist'+workId).attr("view")
            //如果第一次展开评论，通过ajax到后读取
            if(hasComment == "0"){
              $.get("/pro/comment?t="+(new Date()).valueOf() ,{"pid":workId},function(data){
                  var json = JSON.parse(data)
                  for(var i=0; i<json.length; i++){
                      $("#commlist"+workId).append('<div class="media"><a class="pull-left" href="#"><img class="media-object img-rounded" src="'+json[i].Reviewer.ProfileImg+'"></a><div class="media-body"><h6 class="media-heading">'+json[i].Reviewer.UserName+' <span class="commentTime">'+new Date(json[i].PostTime).format("yyyy/MM/dd hh:mm")+'</span></h6><p>'+json[i].CommentDesc+'</p></div></div>');
                  }
                  $('#commlist'+workId).attr("view","1")
              });
            }

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

        //发表评论
        $("#sendCommnet").click(function(){
          var proid = $(this).attr("proid");
          var commdesc = $(".commentdesc"+proid).val();

          $.post("/pro/comment",
            {"commentdesc":commdesc,"proid":proid},
            function(data){
              var j = JSON.parse(data)
              if(j.StateCode == -1){
                alert(j.StateInfo);
              } else {
                $(".commentdesc"+proid).val('');
                $("#commlist"+proid).prepend('<div class="media"><a class="pull-left" href="#"><img class="media-object img-rounded" src="'+j.Reviewer.ProfileImg+'"></a><div class="media-body"><h6 class="media-heading">'+j.Reviewer.UserName+' <span class="commentTime">'+new Date(j.PostTime).format("yyyy/MM/dd hh:mm")+'</span></h6><p>'+j.CommentDesc+'</p></div></div>');
                $(".commnum"+proid).text(parseInt($(".commnum"+proid).text())+1)
              }
            });
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