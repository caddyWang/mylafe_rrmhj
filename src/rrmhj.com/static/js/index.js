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

        //锚动画
        $(document).ready(function(){
          $(".arrow").anchorGoWhere({target:1});
        });

        //登录模态框
        $('#myModal').modal({show:false})

        //指向作品出现收藏按钮，离开则隐藏
        $(".product").bind("mouseover",function(){
          $(this).children(".like").show();
          return false;
        });
        $(".product").bind("mouseleave",function(){
          $(this).children(".like").hide();
          return false;
        });

        //顶 调用相关的表情选择框
        $(".ding").click(function(){
          var uid = $(this).attr("data-uid");
          $('#face_'+uid).fadeIn();
        });
        //关闭表情选择框
        $(".faceclose").click(function(){
          var uid = $(this).attr("data-uid");
          $('#face_'+uid).hide();
        });

        //分享 调用相关的平台选择框
        $(".share").click(function(){
          var uid = $(this).attr("data-uid");
          $('#share_'+uid).fadeIn();
        });
        //关闭分享平台选择框
        $(".shareclose").click(function(){
          var uid = $(this).attr("data-uid");
          $('#share_'+uid).hide();
        });


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
          var workId = $(this).attr("data-uid");
          if($('#comments_'+workId).is(':hidden')) {
            var hasComment = $('#commlist'+workId).attr("view")
            //如果第一次展开评论，通过ajax到后台读取
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
          'contentPage': '/', 
          'scrollTarget': $(window), 
          'heightOffset': 5, 
          'beforeLoad': function(){ 
            $('#loading').fadeIn(); 


            var pageSize = parseInt($('#pageSize').val());
            var proCount = parseInt($('#proCount').val());

            var a = $('#container').children('.thumbnails').size();
            var pageIndex =Math.floor(a / pageSize)

            if(pageIndex*pageSize >= proCount) {
              $('#loading').fadeOut();
              $('#container').stopScrollPagination();
              
              $('#container').append('<div class="alert alert-success" style="text-align:center; display:none;"><button type="button" class="close" data-dismiss="alert">&times;</button><strong>没有新的作品了，等大家来创作吧...</strong></div>')
              $('.alert').delay(1000).fadeIn(0);
            }
            
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


        //分享公共平台地址调用
        $(".bds_tsina").click(function(){
          var img = $(this).attr("img");
          var info = $(this).attr("info");

          var url = "http://service.weibo.com/share/share.php?url=&appkey=3269145958&title="+encodeURIComponent(info)+"&pic="+encodeURIComponent(img)+"&ralateUid=3125160187";
          window.open(url);
        });
        $(".bds_tqq").click(function(){
          var img = $(this).attr("img");
          var info = $(this).attr("info");

          var url = "http://share.v.t.qq.com/index.php?c=share&a=index&url="+encodeURIComponent(img)+"&appkey=801378372&title="+encodeURIComponent(info)+"&pic="+encodeURIComponent(img)+"&line1=";
          window.open(url);
        });
        $(".bds_tqzone").click(function(){
          var img = $(this).attr("img");
          var info = $(this).attr("info");

          var url = "http://sns.qzone.qq.com/cgi-bin/qzshare/cgi_qzshare_onekey?url=http://renrenmanhua.com&showcount=0&summary=&title="+encodeURIComponent(info)+"&site="+encodeURIComponent("人人漫画家")+"&pics="+encodeURIComponent(img)+"&style=103&width=71&height=22&otype=share";
          window.open(url);
        });
        $(".bds_trenren").click(function(){
          var img = $(this).attr("img");
          var info = $(this).attr("info");

          var url = "http://widget.renren.com/dialog/share?resourceUrl=http://renrenmanhua.com&pic="+encodeURIComponent(img)+"&title="+encodeURIComponent("人人漫画家")+"&description="+encodeURIComponent(info)+"&images="+encodeURIComponent(img)+"&charset=utf-8";
          window.open(url);
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
        up.addClass("btn-warning")
      }
      else { 
        var num = down.find(".num");
        num.text(parseInt(num.text())+optView);
        playPlus(down, optView); 
        down.addClass("btn-warning")
      }

      up.addClass("disabled")
      down.addClass("disabled")
      up.unbind("click")
      down.unbind("click")

      $.get("/pro/updown?t="+(new Date()).valueOf() ,{"proId":workId, "optView":optView});
    }


    jQuery.fn.anchorGoWhere = function(options){
     var obj = jQuery(this);
     var defaults = {target:0, timer:1000};
     var o = jQuery.extend(defaults,options);
     obj.each(function(i){
         jQuery(obj[i]).click(function(){
             var _rel = jQuery(this).attr("href").substr(1);
             switch(o.target){
                 case 1: 
                     var _targetTop = jQuery("#"+_rel).offset().top;
                     jQuery("html,body").animate({scrollTop:_targetTop},o.timer);
                     break;
                 case 2:
                     var _targetLeft = jQuery("#"+_rel).offset().left;
                     jQuery("html,body").animate({scrollLeft:_targetLeft},o.timer);
                     break;
             }
             return false;
        });                  
     });
   };