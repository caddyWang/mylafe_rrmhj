<!DOCTYPE html>
{{ $sfu := .SFUrl }}
{{ $pfu := .PFUrl }}

<script src="{{$sfu}}/js/jquery.js"></script>
<script src="{{$sfu}}/js/common.js"></script>

  <script>
    $(function(){

    	if (browser.versions.ios && (browser.versions.iPhone || browser.versions.iPad)) {
    		window.location.href="https://itunes.apple.com/cn/app/ren-ren-man-hua-jia/id608827447?mt=8"
    	}else {
    		window.location.href="{{$pfu}}/littlecartoonist.apk"
    	}
    });
  </script>