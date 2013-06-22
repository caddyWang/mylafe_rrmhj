
{{ $sfu := .SFUrl }}
{{ $isLogin := .IsLogin}}

			{{with .Plist}}
			{{range .}}
            <ul class="thumbnails">
              <li>
                	<div class="thumbnail">
                  		<img src="{{$sfu}}/{{.ImgPath}}" alt="{{.Desc | html2str}}">
                  		<div class="caption">
                    		<span><img src="{{.Author.ProfileImg | fmtHeadImg}}" class="img-rounded"></span> <span>{{.Author.UserName}}</span> <span>{{date .PostTime "Y-m-d" }}</span>
                        <p>{{.Desc | html2str}}</p>
                    	<div class="row-fluid">
                          	<div name="spanBolder"><a href="javascript:void(0);" class="btn btn-small {{.UpNumScript}}"><i class="icon-thumbs-up"></i> <i class="num">{{.UpNum}}</i></a> <a href="javascript:void(0);" class="btn btn-small {{.DownNumScript}}"><i class="icon-thumbs-down"></i> <i class="num">{{.DownNum}}</i></a><input type="hidden" class="uid" value="{{.Pid}}"></span>
	                          <div name="spanBolder" class="btn-group pull-right">
	                            <a href="javascript:void(0);" class="btn btn-small comment"><i class="icon-comment"></i> <i class="commnum{{.Pid}}">{{.CommentNum}}</i></a> 
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

                        <div class="comment_login" {{$isLogin | logoutDisplay}}>发布评论要登录哦：<a href="#qqLogin" class="btn btn-small btn-info" role="button" data-toggle="modal">用腾讯QQ登录</a> <a class="btn btn-small btn-danger" href="https://api.weibo.com/oauth2/authorize?client_id=3269145958&response_type=code&redirect_uri=127.0.0.1:8080/sinalogin/" target="_blank">用新浪微博登录</a></div>
                        <div class="comment_input" {{$isLogin | loginDisplay}}><span><textarea class="commentdesc{{.Pid}}" placeholder="我也来说点什么呗..." rows="1" ></textarea></span> <span class="pull-right"><button class="btn btn-large" id="sendCommnet" type="button" proid="{{.Pid}}">发布</button></span></div>

                        <div class="caption" id="commlist{{.Pid}}" view="0"></div>
                      </div>

                  	</div>
                 </div>
              </li>
            </ul>
          <div class="blank"></div>
          {{end}}
          {{end}}
