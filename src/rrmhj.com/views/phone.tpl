<!DOCTYPE html>

{{ $sfu := .SFUrl }}
{{ $isLogin := .IsLogin}}
{{ $sinaLogin := .SinaLogin}}
{{ $tencLogin := .TencLogin}}

<html lang="zh">
	<head>
		<meta charset="utf-8">
		<title>APP下载 - 人人漫画家</title>
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
    	<meta name="description" content="">
    	<meta name="author" content="">

    	<link href="{{$sfu}}/css/bootstrap.min.css" rel="stylesheet">
      <link href="{{$sfu}}/css/bootstrap-responsive.min.css" rel="stylesheet">
      <link href="{{$sfu}}/css/headerfooter.css" rel="stylesheet">
      <link href="{{$sfu}}/css/prolist.css" rel="stylesheet">
      <link href="{{$sfu}}/css/down.css" rel="stylesheet">

    	<!-- HTML5 shim, for IE6-8 support of HTML5 elements -->
	    <!--[if lt IE 9]>
	      <script src="js/html5shiv.js"></script>
	    <![endif]-->

    	
	</head>

	<body id="top">

		<div class="navbar">
			<div class="navbar-inner">
        <div id="head" class="container-fluid">
            <div class="headleft"><div id="logo"></div></div>
            <div class="headright pull-right dropdown">
                <div class="topmenu-index" onclick="javascript:location.href='/'">首页</div>
                <div class="topmenu-phone"onclick="javascript:location.href='/phone'">软件下载</div>
                <div class="topmenu-login dropdown-toggle" data-close-others="true" data-hover="dropdown" data-toggle="modal" href="#myModal" {{$isLogin | logoutDisplay}}>登录</div>
                <div id="logined" class="topmenu-logined dropdown-toggle" data-close-others="true" data-hover="dropdown" {{$isLogin | loginDisplay}}>{{.UserName}}<div class="arrow"></div></div>
                <ul class="dropdown-menu">
                    <div class="dropdown-arrow"></div>
                    <li {{$isLogin | logoutDisplay}}><a title='新浪微博登录' href="{{$sinaLogin}}" target="_blank"><div class="sinaweibo"></div> 新浪微博登录</a></li>
                    <li {{$isLogin | logoutDisplay}}><a title='腾讯微博登录' href="{{$tencLogin}}" target="_blank"><div class="tencweibo"></div> 腾讯微博登录</a></li>
                    <li {{$isLogin | loginDisplay}}><a title='我的作品' href="/my/pro"><div class="myproduct"></div> 我的作品</a></li>
                    <li {{$isLogin | loginDisplay}}><a title='我的收藏' href="/my/like"><div class="star"></div> 我的收藏</a></li>
                    <li {{$isLogin | loginDisplay}}><a title='退出' href="/my/logout"><div class="exit"></div> 退出</a></li>
                </ul>
              </div>
        </div>
      </div>

      <div class="banner">
        <div class="down-info">
          <div class="down-left"></div>
          <div class="down-right">
            <h1>最简单的创作<br>与分享漫画的方式</h1>
            <h5>专为喜欢漫画的人定制的设计工具。一拖一拉即可用漫画记录你的生活。</h5>
            <div class="down-btn">
              <div class="ios"><div class="icon-ios"></div>iphone版</div>
              <div class="android dropdown-toggle" data-close-others="true" data-hover="dropdown"><div class="icon-android"></div>Android版 <div class="arrow"></div></div>
                <ul class="dropdown-menu downlist">
                    <div class="dropdown-arrow"></div>
                    <li><a title='官方本地下载' href="http://m.mylafe.cc/littlecartoonist1.2.5.apk" target="_blank">官方本地下载</a></li>
                    <li><a title='小米应用下载' href="http://app.xiaomi.com/detail/37624" target="_blank">小米应用下载</a></li>
                    <li><a title='豌豆夹' href="http://www.wandoujia.com/apps/com.shane.littlecartoonist" target="_blank">豌豆夹</a></li>
                    <li><a title='360手机助手' href="http://zhushou.360.cn/detail/index/soft_id/360268" target="_blank">360手机助手</a></li>
                </ul>
            </div>
          </div>
        </div>
      </div>
		</div>

   <div class="introduction">
      <div class="info1">
        <h4>最简单的方式</h4>
        <div>只需要一拖一拉就能得到你想要的故事人物和场景。</div>
        <div class="icon-info1"></div>
      </div>
      <div class="info2">
        <h4>丰富的素材</h4>
        <div>超过100种的特色表情和动作可以自由变换不同组合的情景。</div>
        <div class="icon-info2"></div>
      </div>
      <div class="info3">
        <h4>自由的分享</h4>
        <div>所有这些都可一键分享到微信，新浪微博，腾讯微博，QQ空间等社区。</div>
        <div class="icon-info3"></div>
      </div>
    </div>

  <div class="bottom">
      <div>关于我们  &nbsp;&nbsp;&nbsp;&nbsp;手机客户端  &nbsp;&nbsp;&nbsp;&nbsp;免责声明</div>
      <div>@renrenmanhua.com</div>
  </div>


  <script src="{{$sfu}}/js/jquery.js"></script>
  <script src="{{$sfu}}/js/bootstrap.min.js"></script>
  <script src="{{$sfu}}/js/twitter-bootstrap-hover-dropdown.min.js"></script>

  <script>
    $(function(){
      $(".ios").click(function(){
        window.open("https://itunes.apple.com/cn/app/ren-ren-man-hua-jia/id608827447?mt=8");
      });
      $(".android").click(function(){
        window.open("http://m.mylafe.cc/littlecartoonist1.2.5.apk");
      });
    });
  </script>


  </body>
</html>
