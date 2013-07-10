<!DOCTYPE html>
{{ $sfu := .SFUrl }}
{{ $isLogin := .IsLogin}}
{{ $sinaLogin := .SinaLogin}}
{{ $tencLogin := .TencLogin}}
<html lang="zh">
	<head>
		<meta charset="utf-8">
		<title>人人漫画家</title>
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
    	<meta name="description" content="">
    	<meta name="author" content="">

    	<link href="{{$sfu}}/css/bootstrap.min.css" rel="stylesheet">
    	<link href="{{$sfu}}/css/bootstrap-responsive.min.css" rel="stylesheet">
    	<link href="{{$sfu}}/css/index.css" rel="stylesheet">

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
                <div id="login" class="dropdown-toggle" data-close-others="true" data-hover="dropdown" data-toggle="modal" href="#myModal" {{$isLogin | logoutDisplay}}></div>
                <div id="logined" class="dropdown-toggle" data-close-others="true" data-hover="dropdown" {{$isLogin | loginDisplay}}>{{.UserName}}<div class="arrow"></div></div>
                <ul class="dropdown-menu">
                    <div class="dropdown-arrow"></div>
                    <li {{$isLogin | logoutDisplay}}><a title='新浪微博登录' href="{{$sinaLogin}}" target="_blank"><div class="sinaweibo"></div> 新浪微博登录</a></li>
                    <li {{$isLogin | logoutDisplay}}><a title='腾讯微博登录' href="{{$tencLogin}}" target="_blank"><div class="tencweibo"></div> 腾讯微博登录</a></li>
                    <li {{$isLogin | loginDisplay}}><a title='我的作品' href="#"><div class="myproduct"></div> 我的作品</a></li>
                    <li {{$isLogin | loginDisplay}}><a title='我的收藏' href="#"><div class="star"></div> 我的收藏</a></li>
                    <li {{$isLogin | loginDisplay}}><a title='退出' href="#"><div class="exit"></div> 退出</a></li>
                </ul>
              </div>
        </div>
      </div>

      <div id="ad" class="hidden-phone">
        <div id="default_ad" class="container-fluid">
          <div class="headleft"><img src="{{$sfu}}/img/banner.png"></div>
          <div class="headright">
            <h1>最简单的创作<br>与分享漫画的方式</h1>
            <h5>专为喜欢漫画的人定制的设计工具。一拖一拉即可用漫画记录你的生活。</h5>
            <a class="btn btn-large btn-info">了解详情</a>
          </div>
        </div>
        <div href="#container" class="arrow"></div>
      </div>
    </div>


		<div id="container" class="container-fluid">
      <input type="hidden" id="pageIndex" pageindex="{{.PageIndex}}">
      <input type="hidden" id="proCount" value="{{.ProCount}}">
      <input type="hidden" id="pageSize" value="{{.PageSize}}">

			{{with .Plist}}
			{{range .}}
        <ul class="thumbnails">
              <li>
                  <div class="thumbnail">
                      <div class="product">
                        <div class="like"><div class="star"></div>收藏</div>
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
      
                        <div class="comment_login" {{$isLogin | logoutDisplay}}>发布评论要登录哦：
                          <div class="btnSinaWeiboMini" data-url="{{$sinaLogin}}"><div class="sinaweiboWhite"></div> 新浪微博</div>
                          <div class="btnTencWeiboMini" data-url="{{$tencLogin}}"><div class="tencweiboWhite"></div> 腾讯微博</div>
                        </div>

                        <div class="comment_input" {{$isLogin | loginDisplay}}>
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

  <div class="bottom">
      <div>关于我们  &nbsp;&nbsp;&nbsp;&nbsp;手机客户端  &nbsp;&nbsp;&nbsp;&nbsp;免责声明</div>
      <div>@renrenmanhua.com</div>
  </div>

  <div id="myModal" class="modal hide fade">
        <div class="headleft">
          <h3>欢迎回来</h3>
          <p>目前仅提供新浪和微博的登录。暂不支持其它方式。因为麦麦觉得这样比较简单，比较快咧:)</p>
          <div>
            <div class="btnSinaWeibo" data-url="{{$sinaLogin}}"><div class="sinaweiboWhite"></div> 使用新浪微博登录</div>
            <div class="btnTencWeibo" data-url="{{$tencLogin}}"><div class="tencweiboWhite"></div> 使用腾讯微博登录</div>
          </div>
        </div>
        <div class="headright">
          <div><button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button></div>
          <h6>麦麦提醒你：</h6>
          <p>登录后你可以点评、收藏、分享麦友滴作品，还有好多好涅...</p>
        </div>
      </div> 


	<script src="{{$sfu}}/js/jquery.js"></script>
	<script src="{{$sfu}}/js/bootstrap.min.js"></script>
  <script src="{{$sfu}}/js/twitter-bootstrap-hover-dropdown.min.js"></script>
  <script src="{{$sfu}}/js/scrollpagination.js"></script>
  <script src="{{$sfu}}/js/index.js"></script>


	</body>
</html>
