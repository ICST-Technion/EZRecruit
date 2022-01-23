// API Reference: https://www.wix.com/velo/reference/api-overview/introduction
// “Hello, World!” Example: https://learn-code.wix.com/en/article/1-hello-world
import {local} from 'wix-storage';
import wixWindow from 'wix-window';

$w.onReady(function () {
	let job = local.getItem("job");
	let user = makeid(10)
	let email = "feelfree@tochange.me"
	let phone = "123456789"
	let labels = "programmer,TA,technion,haifa"
	$w('#html1').postMessage("job:"+job)
	$w('#html1').postMessage("user:"+user)
	$w('#html1').postMessage("status:חדש")
	$w('#html1').postMessage("email:"+email)
	$w('#html1').postMessage("phone:"+phone)
	$w('#html1').postMessage("labels:"+labels)
});

function makeid(length) {
    var result           = '';
    var characters       = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    var charactersLength = characters.length;
    for ( var i = 0; i < length; i++ ) {
      result += characters.charAt(Math.floor(Math.random() * charactersLength));
   }
   return result;
}