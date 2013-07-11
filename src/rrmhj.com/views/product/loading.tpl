
{{ $sfu := .SFUrl }}
{{ $isLogin := .IsLogin}}
{{ $sinaLogin := .SinaLogin}}
{{ $tencLogin := .TencLogin}}
{{ $uid := .Uid}}

			{{with .Plist}}
			{{range .}}
            <ul class="thumbnails">
              <li>
                  <div class="thumbnail">
                      <div class="product">
                        <div class="like {{islike $uid .Pid}}" data-pid="{{.Pid}}" data-login="{{$isLogin}}">{{displayLike $uid .Pid}}</div>
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
