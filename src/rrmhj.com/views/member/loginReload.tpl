{{ $sfu := .SFUrl }}

<script src="{{$sfu}}/js/jquery.js"></script>
<script>
	$(function(){
		opener.location.reload();
		window.close();
	});
</script>