<!DOCTYPE html>
{{ $sfu := .SFUrl }}
{{ $isLogin := .IsLogin}}
{{ $sinaLogin := .SinaLogin}}
{{ $tencLogin := .TencLogin}}
<html lang="zh">
  <head>
    <meta charset="utf-8">
    <title>上传作品</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="">
    <meta name="author" content="">

    <!-- CSS -->
    <link href="{{$sfu}}/css/bootstrap.min.css" rel="stylesheet">
    <style type="text/css">

      /* Sticky footer styles
      -------------------------------------------------- */

      html,
      body {
        height: 100%;
        /* The html and body elements cannot have any padding or margin. */
      }

      /* Wrapper for page content to push down footer */
      #wrap {
        min-height: 100%;
        height: auto !important;
        height: 100%;
        /* Negative indent footer by it's height */
        margin: 0 auto -60px;
      }

      /* Set the fixed height of the footer here */
      #push,
      #footer {
        height: 60px;
      }
      #footer {
        background-color: #f5f5f5;
      }

      /* Lastly, apply responsive CSS fixes as necessary */
      @media (max-width: 767px) {
        #footer {
          margin-left: -20px;
          margin-right: -20px;
          padding-left: 20px;
          padding-right: 20px;
        }
      }



      /* Custom page CSS
      -------------------------------------------------- */
      /* Not required for template or sticky footer method. */

      .container {
        width: auto;
        max-width: 680px;
      }
      .container .credit {
        margin: 20px 0;
      }

    </style>
    <link href="{{$sfu}}/css/bootstrap-responsive.min.css" rel="stylesheet">
    <link href="{{$sfu}}/css/headerfooter.css" rel="stylesheet">
    <link href="{{$sfu}}/css/prolist.css" rel="stylesheet">
    <link href="{{$sfu}}/css/index.css" rel="stylesheet">
    <link href="{{$sfu}}/components/uploadify/uploadify.css" rel="stylesheet" type="text/css" />

    <!-- HTML5 shim, for IE6-8 support of HTML5 elements -->
    <!--[if lt IE 9]>
      <script src="../assets/js/html5shiv.js"></script>
    <![endif]-->
  </head>

  <body>


    <!-- Part 1: Wrap all page content here -->
    <div id="wrap">

      <!-- Begin page content -->
      <div class="container">
        <div class="page-header">
          <h2>上传作品</h2>
        </div>
        
        <div class="comment_login" {{$isLogin | logoutDisplay}}>
          <p class="lead">上传作品前请先登录.</p>
          <div class="btnSinaWeiboMini" data-url="{{$sinaLogin}}"><div class="sinaweiboWhite"></div> 新浪微博</div>
          <div class="btnTencWeiboMini" data-url="{{$tencLogin}}"><div class="tencweiboWhite"></div> 腾讯微博</div>
        </div>

        <div {{$isLogin | loginDisplay}}>
          <p class="lead"><img src="{{.UserImg}}" style="width:50px; height:50px; vertical-align: middle;"> {{.UserName}}[{{.Platform}}] <a href="/my/logout?returnurl=/send/pro" style="font-size:12px;">退出</a></p>
          <form class="well form-inline" action="/send/pro" method="post" data-img="">
            <div style="padding-left:20px;">
              <label class="control-label" for="fileInput">选择作品文件：</label> 
              <input type="file" name="detectFile" id="detectFile" multiple="true"/>

              
            </div>

              <div class="comment_input" >
                <input type="hidden" id="imgName" name="imgName" value="{{.Key}}.png">
                <span><textarea class="comm_input" name="descript" placeholder="给作品加个介绍..." rows="3" style="height:50px;" data-pid="0"></textarea></span> <span class="pull-right"><div id="sendCommnet_0">发布</div></span>
              </div>
            
          </form>
          </div>
      </div>

      <div id="push"></div>
    </div>




  <script src="{{$sfu}}/js/jquery.js"></script>
  <script src="{{$sfu}}/js/bootstrap.min.js"></script>
  <script src="{{$sfu}}/js/prolist.js"></script>
  <script src="{{$sfu}}/components/uploadify/jquery.uploadify-3.1.min.js"></script>

  </body>

  <script>
    $(function(){
      //文件上传
      $("#detectFile").uploadify({

        'fileObjName'   :  'file',
        'uploader'       : 'http://up.qiniu.com/',
        'method'       : 'post',
        'formData'     : {'key':'{{.Key}}.png','token':'{{.Uptoken}}'},
        'swf'           : '{{$sfu}}/components/uploadify/uploadify.swf',
        'multi'          : false,
        'fileTypeExts'   : '*.png; *.jpg; *.jpeg; *.gif',
        'fileTypeDesc'   : '只支持图片文件',
        'buttonText'   : '上传文件',
        //'buttonImage'  : '/style-components/uploadify/file.gif',
        'fileSizeLimit'  : '5MB',
        'width'      : 80,
        'heigth'     : 280,
        'onUploadSuccess'  : function(file, data, response){
              var jsonText = eval('('+data+')'); 
              $(".well").append('<a href="http://rrmhj.qiniudn.com/{{.Key}}.png" target="_blank"><img src="http://rrmhj.qiniudn.com/{{.Key}}.png"></a>');
              $(".well").attr("data-img",jsonText.key);

              return false;
              },
        'onUploadError' : function(file, errorCode, errorMsg, errorString){
              alert('文件上传出错！原因：'+errorString);
              }
      });

      $("#sendCommnet_0").click(function(){
        var desc = $(".comm_input").val();
        var img = $(".well").attr("data-img");

        if(img == "") {
          alert('先上传作品！');
          return false;
        } else if(desc == "") {
          alert('多少说点啥呗！');
          return false;
        } else {
          $(".well").submit();
        }

      });
    });
    
  </script>
</html>
