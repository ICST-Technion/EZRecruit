// API Reference: https://www.wix.com/velo/reference/api-overview/introduction
// “Hello, World!” Example: https://learn-code.wix.com/en/article/1-hello-world
import {insertJob} from 'backend/Modules/Jobs/insertJob';
import wixWindow from 'wix-window';

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
export function submitJob_click(event) {
	// This function was added from the Properties & Events panel. To learn more, visit http://wix.to/UcBnC-4
	// Add your code for this event here:
	insertJob($w("#jobNameInput").value, $w("#locationInput").value, $w("#descriptionInput").value,
				$w("#requiredSkillsInput").value, $w("#labelsInput").value, $w("#googleFormsInput").value)
    .then(httpResponse => {
        console.log("httpResponse: " + httpResponse);
    })
    .catch(error => {
        console.log(error);
    });
}

/**
 *	Adds an event handler that fires when a visitor submits a Wix Form and it is successfully received by the server.
 */
export function wixForms1_wixFormSubmitted() {
	// This function was added from the Properties & Events panel. To learn more, visit http://wix.to/UcBnC-4
	// Add your code for this event here:
	wixWindow.lightbox.close()
}