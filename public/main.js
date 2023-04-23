let picked = false
let pickedIndex = -1
document.addEventListener('DOMContentLoaded', function() {
  const images = document.querySelectorAll('.img');
  const pickButton = document.querySelector('.pick-button');
  const fullscreenOverlay = document.querySelector('.fullscreen-overlay');
  const fullscreenImageContainer = document.querySelector('.fullscreen-image-container');
  const fullscreenImage = fullscreenImageContainer.querySelector('img');
  const closeButton = fullscreenImageContainer.querySelector('.close-button');

  // handle click events on image containers
	images.forEach(function(img) {
		img.addEventListener('click', function() {
			const imgSrc = img.getAttribute('src');
			fullscreenImage.setAttribute('src', imgSrc);
			fullscreenOverlay.style.display = 'flex';
			fullscreenOverlay.classList.remove('fadeOut');
		});
	});

  // handle click pick button
  // afeter click add pick class to a random image
  pickButton.addEventListener('click', function() {
    if (picked) {
      images[pickedIndex].classList.remove('pick');
      picked = false
      pickedIndex = -1
    }
    const index = Math.floor(Math.random() * images.length)
    const randomImage = images[index];
    randomImage.classList.add('pick');
    pickedIndex = index
    picked = true


    // scroll to the picked image
    randomImage.scrollIntoView({
      behavior: 'smooth',
      block: 'start',
      inline: 'center'
    });
  });


  // handle click event on close button
  closeButton.addEventListener('click', function() {
    fullscreenOverlay.classList.add('fadeOut');
    setTimeout(function() {
      fullscreenOverlay.style.display = 'none';
    }, 330);
  });});
