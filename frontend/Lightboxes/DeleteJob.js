// API Reference: https://www.wix.com/velo/reference/api-overview/introduction
// “Hello, World!” Example: https://learn-code.wix.com/en/article/1-hello-world
import {local} from 'wix-storage';
import wixWindow from 'wix-window';

import {deleteJob} from 'backend/Modules/Jobs/deleteJob';

$w.onReady(function () {
	// Write your JavaScript here

	// To select an element by ID use: $w('#elementID')

	// Click 'Preview' to run your code
});

/**
*	Adds an event handler that runs when the element is clicked.
	[Read more](https://www.wix.com/corvid/reference/$w.ClickableMixin.html#onClick)
*	 @param {$w.MouseEvent} event
*/
export function closeButton_click(event) {
	// This function was added from the Properties & Events panel. To learn more, visit http://wix.to/UcBnC-4
	// Add your code for this event here:
	wixWindow.lightbox.close();
}

/**
*	Adds an event handler that runs when the element is clicked.
	[Read more](https://www.wix.com/corvid/reference/$w.ClickableMixin.html#onClick)
*	 @param {$w.MouseEvent} event
*/
export function deleteJobButton_click(event) {
	// This function was added from the Properties & Events panel. To learn more, visit http://wix.to/UcBnC-4
	// Add your code for this event here:
	let jobId = local.getItem("jobToDelete");
	console.log(jobId)
	deleteJob(jobId).then(() => wixWindow.lightbox.close());
}