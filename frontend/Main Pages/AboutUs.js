// API Reference: https://www.wix.com/velo/reference/api-overview/introduction
// “Hello, World!” Example: https://learn-code.wix.com/en/article/1-hello-world
import wixUsers from 'wix-users';
$w.onReady(function () {
	// Write your JavaScript here

	// To select an element by ID use: $w('#elementID')

	// Click 'Preview' to run your code
	$w('#button16').hide();
	wixUsers.currentUser.getRoles()
	.then((roles) => {
		const currentUser = wixUsers.currentUser;
		if(currentUser.role === "Admin"){
			$w('#button16').show();
		} else {
			$w('#button16').hide();
		}
	});
});