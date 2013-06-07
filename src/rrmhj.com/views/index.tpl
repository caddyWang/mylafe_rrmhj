<!DOCTYPE html>
{{ $sfu := .SFUrl }}
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

	<body>

		<div class="navbar">
			<div class="navbar-inner">
				<div class="container-fluid" style="text-align:center;">
            <span id="logo-left"></span>
            <img src="{{$sfu}}/img/logo.png">
            <span id="top-nav" class="pull-right"><a href="#"><img id="login"src="{{$sfu}}/img/login.png"></a> <a href="#"><img id="phone" src="{{$sfu}}/img/phone.png" class="hidden-phone"></a></span>
				</div>
			</div>
		</div>

		<div id="container" class="container-fluid">

			{{with .Plist}}
			{{range .}}
            <ul class="thumbnails">
              <li>
                	<div class="thumbnail">
                  		<img src="{{$sfu}}/{{.ImgPath}}" alt="">
                  		<div class="caption">
                    		<span><img src="{{$sfu}}/{{.Author.HeadImg | fmtHeadImg}}" class="img-rounded"></span> <span>dadairen</span> <span>{{date .PostTime "Y-m-d" }}</span>
                        <p>{{.Desc}}</p>
                    	<div class="row-fluid">
                          	<div name="spanBolder"><a href="javascript:void(0);" class="btn btn-small up"><i class="icon-thumbs-up"></i> <i class="num">{{.UpNum}}</i></a> <a href="javascript:void(0);" class="btn btn-small down"><i class="icon-thumbs-down"></i> <i class="num">{{.DownNum}}</i></a><input type="hidden" class="uid" value="{{.Pid}}"></span>
	                          <div name="spanBolder" class="btn-group pull-right">
	                            <a href="javascript:void(0);" class="btn btn-small comment"><i class="icon-comment"></i> {{.Comments | len}}</a> 
	                            <a class="btn btn-small dropdown-toggle" data-delay="1000" data-hover="dropdown" data-toggle="dropdown"><i class="icon-share"></i> 分享</a>
	                            <ul class="dropdown-menu">
	                              <li><a class='bds_tsina' title='分享到新浪微博' href="#" style="padding-left:25px;"> 新浪微博</a></li>
	                              <li><a class='bds_tqq' title='分享到腾讯微博' href="#" style="padding-left:25px;"> 腾讯微博</a></li>
	                              <li><a class='bds_tqzone' title='分享到QQ空间' href='#' style="padding-left:25px;"> QQ空间</a></li>
	                              <li><a class='bds_trenren' title='分享到人人网' href='#' style="padding-left:25px;"> 人人网</a></li>
	                            </ul>
	                            <input type="hidden" class="uid" value="{{.Pid}}">
	                          </div>
	                        </div>
                  		</div>

                      <div id="comments_{{.Pid}}" style="display:none">
                        <p class="line"></p>

                        <div class="caption">
                          {{with .Comments}}
                          {{range .}}
                          <div class="media">
                            <a class="pull-left" href="#"><img class="media-object img-rounded" src="{{$sfu}}/{{.Reviewer.HeadImg | fmtHeadImg}}"></a>
                            <div class="media-body">
                              <h6 class="media-heading">{{.Reviewer.UserName}}</h6>
                              <p>{{.CommentDesc}}</p>
                            </div>
                          </div>
                          {{end}}
                          {{end}}
                        </div>
                      </div>

                  	</div>
                 </div>
              </li>
            </ul>
          <div class="blank"></div>
          {{end}}
          {{end}}

	    </div>


    <div id="loading"><img src="{{$sfu}}/img/loading.gif"></div>

	<script src="{{$sfu}}/js/jquery.js"></script>
	<script src="{{$sfu}}/js/bootstrap.min.js"></script>
  	<script src="{{$sfu}}/js/twitter-bootstrap-hover-dropdown.min.js"></script>
  	<script src="{{$sfu}}/js/scrollpagination.js"></script>
  	<script src="{{$sfu}}/js/index.js"></script>

	</body>
</html>