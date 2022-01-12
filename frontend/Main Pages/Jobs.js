// API Reference: https://www.wix.com/velo/reference/api-overview/introduction
// “Hello, World!” Example: https://learn-code.wix.com/en/article/1-hello-world
import {getJobs} from 'backend/Modules/Jobs/getJobs';
import wixUsers from 'wix-users';
import wixWindow from 'wix-window';
import {local} from 'wix-storage';

$w.onReady(function () {
	// admin user perms
	//setTypeOfInput();
	wixUsers.currentUser.getRoles()
	.then((roles) => {
		const currentUser = wixUsers.currentUser;
		if(currentUser.role === "Admin"){
			$w('#button16').hide();
		} else {
			$w('#button16').show();
		}
	});
	// set jobsRepeater callback
	$w("#jobsRepeater").onItemReady( ($item, itemData, index) => {
    		$item("#jobTitle").text = itemData.title;
    		$item("#jobDescription").text = itemData.description;
			$item("#jobLocation").text = itemData.location;
			$item("#labels").text = itemData.labels.toString();
			$item("#jobId").text = itemData._id;
			$item("#chooseJobButton").link = itemData.formLink;
  	});

	refreshJobs()

  //puts the job options in the dropdown menu
  getJobs().then(jobInfo => {
    let jobsInfo = jobInfo
    let jobsDropdownOptions = []
    for(let i in jobsInfo) {
      let option = {"label": jobsInfo[i].title, "value": jobsInfo[i]._id}
      jobsDropdownOptions.push(option);
    }
    $w("#jobDropDown").options = jobsDropdownOptions
	})

	$w("#jobsPagination").currentPage = 1;


});



export function refreshJobs() {
	// fetch jobs into repeater
	getJobs().then(jobInfo => {
    	$w("#jobsRepeater").data = []
		$w("#jobsRepeater").data = jobInfo.value;
	})
}

/**
*	Adds an event handler that runs when the element is clicked.
	[Read more](https://www.wix.com/corvid/reference/$w.ClickableMixin.html#onClick)
*	 @param {$w.MouseEvent} event
*/
export function deleteJob_click(event, $w) {
	// This function was added from the Properties & Events panel. To learn more, visit http://wix.to/UcBnC-4
	// Add your code for this event here:
	local.setItem("jobToDelete", $w("#jobId").text)
  console.log($w("#jobId").text)
	wixWindow.openLightbox("DeleteJob").then(() => {
		refreshJobs();
	});
}

/**
*	Adds an event handler that runs when the element is clicked.
	[Read more](https://www.wix.com/corvid/reference/$w.ClickableMixin.html#onClick)
*	 @param {$w.MouseEvent} event
*/
export function AddJobButton_click(event) {
	// This function was added from the Properties & Events panel. To learn more, visit http://wix.to/UcBnC-4
	// Add your code for this event here:
	wixWindow.openLightbox("AddJob").then(() => {
		refreshJobs();
	});
}

/**
*	Adds an event handler that runs when the element is clicked.
	[Read more](https://www.wix.com/corvid/reference/$w.ClickableMixin.html#onClick)
*	 @param {$w.MouseEvent} event
*/
export function searchJobsButton_click(event) {
	// This function was added from the Properties & Events panel. To learn more, visit http://wix.to/UcBnC-4
	// Add your code for this event here:
	getJobs($w("#searchJobInput").value).then(jobInfo => {
    	$w("#jobsRepeater").data = []
		$w("#jobsRepeater").data = jobInfo;
	})
}

/**
*	Adds an event handler that runs when an input element's value
 is changed.
	[Read more](https://www.wix.com/corvid/reference/$w.ValueMixin.html#onChange)
*	 @param {$w.Event} event
*/
export function searchJobInput_change(event) {
	// This function was added from the Properties & Events panel. To learn more, visit http://wix.to/UcBnC-4
	// Add your code for this event here:
  if ($w("#searchJobInput").value.length == 0){
    refreshJobs()
  }

}

export function jobsPagination_change(event) {
	// This function was added from the Properties & Events panel. To learn more, visit http://wix.to/UcBnC-4
	// Add your code for this event here:
}