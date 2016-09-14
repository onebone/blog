$(function(){
	if(location.protocol !== 'https'){
		$('#warning').html(`<i class="material-icons" style="float: left;">warning</i>Your connection is not over HTTPS.<br>Your username, password may not be private.`);
	}
})
