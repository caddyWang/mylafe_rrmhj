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

        loadPageInit();

        //返回顶部
        $(window).scroll(function(){
          if($('#back-to-top').is(':hidden')){
            $("#back-to-top").css({bottom:"100px"});
            $("#back-to-top").removeClass("fly-to-top");
          }

          if ($(window).scrollTop()>100){
            $("#back-to-top").fadeIn(1000);
          } else {
            $("#back-to-top").fadeOut(800);
          }
        });

        $("#back-to-top").click(function(){
          $("#back-to-top").addClass("fly-to-top");
          $('body,html').animate({scrollTop:0},1000);
          $("#back-to-top").animate({bottom:400},1000);
          return false;
        });

        //锚动画
        $(document).ready(function(){
          $(".arrow").anchorGoWhere({target:1});
        });

        //登录模态框
        $('#myModal').modal({show:false})

    });

    //无限数据读取
    function newScrollData(ajaxURL) {
      var pageIndex = $("#pageIndex").attr("data-val");
      var pageSize = parseInt($('#pageSize').val());
      var proCount = parseInt($('#proCount').val());

      if (pageIndex*pageSize >= proCount){
        $('.alert').delay(1000).fadeIn(0);
      } else {
        $.ajax({
          type: 'POST',
          url: ajaxURL,
          data: {"pageIndex":  pageIndex},
          success: function(data){
            $("#pageIndex").attr("data-val", parseInt(pageIndex)+1);
            $("#container").append(data);
            loadPageInit();
          },
          dataType: 'html'
        });

        $('#loading').show();
      }

      $('#loading').fadeOut();

    }


    function loadPageInit() {
      //指向作品出现收藏按钮，离开则隐藏
        $(".product").bind("mouseover",function(){
          $(this).children(".like").show();
          return false;
        });
        $(".product").bind("mouseleave",function(){
          $(this).children(".like").hide();
          return false;
        });

        //收藏
        $(".unlike").unbind("click");
        $(".unlike").click(function(){
          var pid = $(this).attr("data-pid");
          var login = $(this).attr("data-login");
          var like = $(this);
          if(login == "false") {
            $('#myModal').modal('show');
            return false;
          }else{
            $.get("/pro/like?t="+(new Date()).valueOf() ,{"proId":pid},function(data){
                  if(data == "0"){
                    like.html("<div class='star-full'></div>已收藏");
                    like.unbind("click");
                  }
              });
          }
        });

        //删除创作的作品
        $(".delmypro").unbind("click");
        $(".delmypro").click(function(){
          var pid = $(this).attr("data-pid");
          var login = $(this).attr("data-login");
          var like = $(this);
          if(login == "false") {
            $('#myModal').modal('show');
            return false;
          }else{
            $.get("/pro/delpro?t="+(new Date()).valueOf() ,{"proId":pid},function(data){
                  if(data == "0"){
                    $("#thumbnails_"+pid).fadeOut();

                    var w = findScreenW();
                    var leftS = parseInt(w)/2 - 100;
                    $(".deleted-info").css({left : leftS+"px"})
                    $(".deleted-info").show(500).delay(1000).hide(500);

                    var myproNum = parseInt($("#myproNum").text())-1;
                    $("#myproNum").text(myproNum);
                  }else if(data == "-2"){
                    alert("登录超时，请重新登录后再删除！");
                  }else{
                    alert("删除出错！")
                  }
              });
          }
        });

        //顶 调用相关的表情选择框
        $(".up").unbind("click");
        $(".up").click(function(){
          var uid = $(this).attr("data-uid");
          $('#face_'+uid).fadeIn();
        });
        //顶过用户的提醒
        $(".ding_disabled").unbind("click");
        $(".ding_disabled").click(function(){
          var uid = $(this).attr("data-uid");
          $('#has_ding_'+uid).show(300).delay(1000).hide(300);
        });
        //关闭表情选择框
        $(".faceclose").unbind("click");
        $(".faceclose").click(function(){
          var uid = $(this).attr("data-uid");
          $('#face_'+uid).fadeOut();
        });

        //分享 调用相关的平台选择框
        $(".share").unbind("click");
        $(".share").click(function(){
          var uid = $(this).attr("data-uid");
          $('#share_'+uid).fadeIn();
        });
        //关闭分享平台选择框
        $(".shareclose").unbind("click");
        $(".shareclose").click(function(){
          var uid = $(this).attr("data-uid");
          $('#share_'+uid).fadeOut();
        });

        //新浪微博登录绑定
        $(".btnSinaWeiboMini").unbind("click");
        $(".btnSinaWeiboMini").click(function(){
          var durl = $(this).attr("data-url");
          window.open(durl);
        });
        $(".btnSinaWeibo").unbind("click");
        $(".btnSinaWeibo").click(function(){
          var durl = $(this).attr("data-url");
          window.open(durl);
        });
        //腾讯微博登录绑定
        $(".btnTencWeiboMini").unbind("click");
        $(".btnTencWeiboMini").click(function(){
          var durl = $(this).attr("data-url");
          window.open(durl);
        });
        $(".btnTencWeibo").unbind("click");
        $(".btnTencWeibo").click(function(){
          var durl = $(this).attr("data-url");
          window.open(durl);
        });


        //顶
        $(".ding-face").find('div[class|="face"]').unbind("click");
        $(".ding-face").find('div[class|="face"]').click(function(){
          var workId = $(this).attr("data-pid");
          var dingface = $(this).attr("data-val");
          var up = $(".ding_"+workId);
          $('#face_'+workId).fadeOut();
          optUpOrDownOrAttention(up, workId, dingface);
          
        });

        //展现评论
        $(".comment").unbind("click");
        $(".comment").click(function(){
          var workId = $(this).attr("data-uid");
          if($('#comments_'+workId).is(':hidden')) {
            var hasComment = $('#comments_'+workId).attr("view")
            //如果第一次展开评论，通过ajax到后台读取
            if(hasComment == "0"){
              $.get("/pro/comment?t="+(new Date()).valueOf() ,{"pid":workId},function(data){
                  var json = JSON.parse(data)
                  for(var i=0; i<json.length; i++){
                      $(".comm_list_"+workId).append('<div class="comments caption"><div class="user"><span><img class="media-object img-rounded" src="'+json[i].Reviewer.ProfileImg+'"></span><span>'+json[i].Reviewer.UserName+'</span></div><div class="content pull-right"><span>'+json[i].CommentDesc+'</span><span class="time pull-right">'+new Date(json[i].PostTime).format("yyyy/MM/dd")+'</span></div>');
                  }
                  $('#comments_'+workId).attr("view","1")
              });
            }

            $('#comments_'+workId).fadeIn();
          } else {
            $('#comments_'+workId).fadeOut();
          }
        })

        //评论输入框选中
        $(".comm_input").focus(function(){
          var uid = $(this).attr("data-pid");
          $("#sendCommnet_"+uid).addClass("comment_input_focus");
        });
        //评论输入框失去焦点
        $(".comm_input").blur(function(){
          var uid = $(this).attr("data-pid");
          $("#sendCommnet_"+uid).removeClass("comment_input_focus");
        });

        //发表评论
        $(".sendCommnet").unbind("click");
        $(".sendCommnet").click(function(){
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
                $(".comm_list_"+proid).prepend('<div class="comments caption"><div class="user"><span><img class="media-object img-rounded" src="'+j.Reviewer.ProfileImg+'"></span><span>'+j.Reviewer.UserName+'</span></div><div class="content pull-right"><span>'+j.CommentDesc+'</span><span class="time pull-right">'+new Date(j.PostTime).format("yyyy/MM/dd")+'</span></div>');
                $(".commnum"+proid).text(parseInt($(".commnum"+proid).text())+1)
                $("#sendCommnet_"+proid).removeClass("comment_input_focus");
              }
            });
        });


        //分享公共平台地址调用
        $('div[class|="shareicon"]').unbind("click");
        $('div[class|="shareicon"]').click(function(){
          var img = $(this).attr("img");
          var info = $(this).attr("info");
          var uid = $(this).attr("data-uid");
          $('#share_'+uid).hide();

          var url = "";
          if($(this).hasClass("shareicon-sina")) {
              url = "http://service.weibo.com/share/share.php?url=&appkey=3269145958&title="+encodeURIComponent(info)+"&pic="+encodeURIComponent(img)+"&ralateUid=3125160187";
          }else if($(this).hasClass("shareicon-tenc")) {
              url = "http://share.v.t.qq.com/index.php?c=share&a=index&url="+encodeURIComponent(img)+"&appkey=801378372&title="+encodeURIComponent(info)+"&pic="+encodeURIComponent(img)+"&line1=";
          }else if($(this).hasClass("shareicon-qq")) {
              url = "http://sns.qzone.qq.com/cgi-bin/qzshare/cgi_qzshare_onekey?url=http://renrenmanhua.com&showcount=0&summary=&title="+encodeURIComponent(info)+"&site="+encodeURIComponent("人人漫画家")+"&pics="+encodeURIComponent(img)+"&style=103&width=71&height=22&otype=share";
          }else if($(this).hasClass("shareicon-renren")) {
              url = "http://widget.renren.com/dialog/share?resourceUrl=http://renrenmanhua.com&pic="+encodeURIComponent(img)+"&title="+encodeURIComponent("人人漫画家")+"&description="+encodeURIComponent(info)+"&images="+encodeURIComponent(img)+"&charset=utf-8";
          }

          window.open(url);
        });
    }


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
    function optUpOrDownOrAttention(up, workId, dingface){
      
      var num = $(".num"+workId);
      var inum = parseInt(num.text())+1;
      num.text(inum);
      playPlus(up, 1); 

      up.addClass("ding_disabled")
      up.unbind("click")
      $(".ding_"+workId).click(function(){
        $('#has_ding_'+workId).show(300).delay(1000).hide(300);
      });

      $.get("/pro/updown?t="+(new Date()).valueOf() ,{"proId":workId, "dingface":dingface});
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

   //获取浏览器宽度
   function findScreenW()  {
      var winWidth;

      //获取窗口宽度
      if (window.innerWidth)
        winWidth = window.innerWidth;
      else if ((document.body) && (document.body.clientWidth))
        winWidth = document.body.clientWidth;

      return winWidth;
   }