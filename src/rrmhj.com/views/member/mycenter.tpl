<!DOCTYPE html>
{{ $sfu := .SFUrl }}
{{ $isLogin := .IsLogin}}
{{ $sinaLogin := .SinaLogin}}
{{ $tencLogin := .TencLogin}}
{{ $uid := .Uid}}
{{ $mypro := .MyPro}}
<html lang="zh">
	<head>
		<meta charset="utf-8">
		<title>我的{{if $mypro}}作品{{else}}收藏{{end}} - 人人漫画家</title>
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
    	<meta name="description" content="">
    	<meta name="author" content="">

    	<link href="{{$sfu}}/css/bootstrap.min.css" rel="stylesheet">
      <link href="{{$sfu}}/css/bootstrap-responsive.min.css" rel="stylesheet">
      <link href="{{$sfu}}/css/headerfooter.css" rel="stylesheet">
      <link href="{{$sfu}}/css/prolist.css" rel="stylesheet">
      <link href="{{$sfu}}/css/my.css" rel="stylesheet">

    	<!-- HTML5 shim, for IE6-8 support of HTML5 elements -->
	    <!--[if lt IE 9]>
	      <script src="js/html5shiv.js"></script>
	    <![endif]-->

    	
	</head>

	<body id="top">

    <div id="back-to-top" href="#top" class="hidden-phone"></div>

		<div class="navbar">
			<div class="navbar-inner">
        <div id="head" class="container-fluid">
            <div class="headleft"><div id="logo"></div></div>
            <div class="headright pull-right dropdown">
                <div id="logined" class="dropdown-toggle" data-close-others="true" data-hover="dropdown">{{.UserName}}<div class="arrow"></div></div>
                <ul class="dropdown-menu">
                    <div class="dropdown-arrow"></div>
                    <li><a title='我的作品' href="/my/pro"><div class="myproduct"></div> 我的作品</a></li>
                    <li><a title='我的收藏' href="/my/like"><div class="star"></div> 我的收藏</a></li>
                    <li><a title='退出' href="/my/logout"><div class="exit"></div> 退出</a></li>
                </ul>
              </div>
        </div>
      </div>
		</div>

    <div id="container" class="container-fluid">
      <input type="hidden" id="pageIndex" data-val="{{.PageIndex}}">
      <input type="hidden" id="proCount" value="{{.ProCount}}">
      <input type="hidden" id="pageSize" value="{{.PageSize}}">

      <div id="my-nav" class="container-fluid">
        {{if $mypro}}<div class="mypro-arrow"></div>{{end}}{{if .MyLike}}<div class="mylike-arrow"></div>{{end}}
        <div class="my-products"><div class="icon-mypro"></div> 我的作品(<span id="myproNum">{{.ProCount}}</span>)</div>
        <div class="my-likes"><div class="icon-mylike"></div> 我的收藏(<span id="mylikeNum">{{.LikeCount}}</span>)</div>
        <div class="deleted-info">删除成功了！</div>
      </div>

      {{if .ListNull}}
      <div class="nopro"></div>
      <div class="return-url"><a href="/">返回首页</a> &nbsp;&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;&nbsp; <a href="javascript:history.back(-1);">返回上一页</a></div>
      {{end}}

      {{with .Plist}}
      {{range .}}
        <ul class="thumbnails" id="thumbnails_{{.Pid}}">
              <li>
                  <div class="thumbnail">
                      <div class="product">
                        <div class="like {{if $mypro}}delmypro{{else}}delmylike{{end}}" data-pid="{{.Pid}}" data-login="{{$isLogin}}"><div class="delpro"></div> 删除</div>
                        <img src="{{$sfu}}/{{.ImgPath}}" alt="{{.Desc | html2str}}">
                      </div>

                      <div class="row-fluid">
                        <div id="face_{{.Pid}}" class="ding-face">
                          <div class="face-laugh" data-pid="{{.Pid}}" data-val="laugh"></div>
                          <div class="face-love" data-pid="{{.Pid}}" data-val="love"></div>
                          <div class="face-applause" data-pid="{{.Pid}}" data-val="applause"></div>
                          <div class="face-chop" data-pid="{{.Pid}}" data-val="chop"></div>
                          <div class="face-cry" data-pid="{{.Pid}}" data-val="cry"></div>
                          <div class="pull-right"><button type="button" class="close faceclose" data-dismiss="modal" aria-hidden="true" data-uid="{{.Pid}}">&times;</button></div>
                        </div>
                        <div id="share_{{.Pid}}" class="share-icon">
                          <div class="shareicon-sina" img="{{$sfu}}/{{.ImgPath}}" info="{{.Desc | html2str}}" data-uid="{{.Pid}}"></div>
                          <div class="shareicon-tenc" img="{{$sfu}}/{{.ImgPath}}" info="{{.Desc | html2str}}" data-uid="{{.Pid}}"></div>
                          <div class="shareicon-qq" img="{{$sfu}}/{{.ImgPath}}" info="{{.Desc | html2str}}" data-uid="{{.Pid}}"></div>
                          <div class="shareicon-renren" img="{{$sfu}}/{{.ImgPath}}" info="{{.Desc | html2str}}" data-uid="{{.Pid}}"></div>
                          <div class="pull-right"><button type="button" class="close shareclose" data-dismiss="modal" aria-hidden="true" data-uid="{{.Pid}}">&times;</button></div>
                        </div>
                        <div class="pro-user">
                            <div><img src="{{.Author.ProfileImg | fmtHeadImg}}"></div>
                            <div class="pull-right"><div class="user">{{.Author.UserName}}</div> <div class="time">{{date .PostTime "Y-m-d" }}</div></div>
                        </div>
                        <div class="pro-opt pull-right">
                          <div id="has_ding_{{.Pid}}" class="has-ding">你顶过了!</div>
                          <div class="ding ding_{{.Pid}} {{.UpNumScript}}" data-uid="{{.Pid}}"><div class="icn-ding"></div><span class="num{{.Pid}}">{{.UpNum}}</span></div>
                          <div class="comment" data-uid="{{.Pid}}"><div class="icn-comm"></div><span class="commnum{{.Pid}}">{{.CommentNum}}</span></div>
                          <div class="share" data-uid="{{.Pid}}"><div class="icn-share"></div>分享</div>
                        </div>

                        <div class="user-arrow"></div>
                      </div>
                      
                      <div class="pro-info">{{.Desc | html2str}}</div>

                      <div id="comments_{{.Pid}}" class="comment-list" view="0">

                        <div class="comment_input">
                          <span><textarea class="commentdesc{{.Pid}} comm_input" data-pid="{{.Pid}}" placeholder="我来说两句..." rows="1" ></textarea></span> <span class="pull-right"><div id="sendCommnet_{{.Pid}}" class="sendCommnet" proid="{{.Pid}}">发布</div></span>
                        </div>

                        <div class="comm_list_{{.Pid}}"></div>

                      </div>

                  </div>
              </li>
            </ul>
            <div class="blank"></div>

          {{end}}
          {{end}}

   </div>

  <div id="loading"><img src="{{$sfu}}/img/loading.gif"><p>内容加载中...</p></div>
  <div class="alert alert-success" style="text-align:center; display:none;"><button type="button" class="close" data-dismiss="alert">&times;</button><strong>没有作品了，{{if $mypro}}快来创作吧{{else}}有喜欢的再收吧{{end}}...</strong></div><br>

  <div class="bottom">
      <div>关于我们  &nbsp;&nbsp;&nbsp;&nbsp;手机客户端  &nbsp;&nbsp;&nbsp;&nbsp;免责声明</div>
      <div>@renrenmanhua.com</div>
  </div>


  <script src="{{$sfu}}/js/jquery.js"></script>
  <script src="{{$sfu}}/js/bootstrap.min.js"></script>
  <script src="{{$sfu}}/js/twitter-bootstrap-hover-dropdown.min.js"></script>
  <script src="{{$sfu}}/js/prolist.js"></script>

  <script>
    $(function(){
      if ("{{.ListNull}}" != "true") {
        $(function(){
          $(window).scroll(function(){
              if  ($(window).scrollTop() == $(document).height() - $(window).height()){
                if ("{{$mypro}}" == "true") {
                  newScrollData("/my/pro");
                }else{
                  newScrollData("/my/like");
                }
              }
          });
        });
      } else {
        $("#loading").hide();
      }

      $(".my-products").click(function(){
        location.href = "/my/pro";
      });
      $(".my-likes").click(function(){
        location.href = "/my/like";
      });
    });
  </script>


  </body>
</html>
