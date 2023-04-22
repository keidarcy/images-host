let picked = false
let pickedIndex
document.addEventListener('DOMContentLoaded', function() {
  const images = document.querySelectorAll('.js-img');
  const imageDivs = document.querySelectorAll('.img-div');
  const fullscreenOverlay = document.querySelector('.fullscreen-overlay');
  const fullscreenImageContainer = document.querySelector('.fullscreen-image-container');
  const fullscreenImage = fullscreenImageContainer.querySelector('img');
  const closeButton = fullscreenImageContainer.querySelector('.close-button');
  const pickButton = document.querySelector('.magic-button');

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
      imageDivs[pickedIndex].classList.remove('pick');
    }
    const index = Math.floor(Math.random() * imageDivs.length)
    const randomImage = imageDivs[index];
    randomImage.classList.add('pick');
    pickedIndex = index
    picked = true
  });


  // handle click event on close button
  closeButton.addEventListener('click', function() {
    fullscreenOverlay.classList.add('fadeOut');
    setTimeout(function() {
      fullscreenOverlay.style.display = 'none';
    }, 330);
  });});
