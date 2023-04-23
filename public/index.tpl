<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
    <link
      href="https://fonts.googleapis.com/css2?family=Satisfy&display=swap"
      rel="stylesheet"
    />
    <link rel="stylesheet" href="./style.css" />
    <title>{{ .Title}}</title>
  </head>

  <body>
    <h1>What's today's food</h1>
    <button class="pick-button"><span>Pick</span></button>
  <main>
    <div class="img-container">
      {{ range.ImgNames }} 
        <img 
          loading="lazy" 
          class="img js-img" 
          alt="food-{{.}}" 
          src="{{$.ImgUrl}}{{.}}" 
        /> 
      {{ end }}
      {{ range.ImgNames }} 
        <img 
          loading="lazy" 
          class="img" 
          alt="food-{{.}}" 
          src="{{$.ImgUrl}}{{.}}" 
        /> 
      {{ end }}
      {{ range.ImgNames }} 
        <img 
          loading="lazy" 
          class="img" 
          alt="food-{{.}}" 
          src="{{$.ImgUrl}}{{.}}" 
        /> 
      {{ end }}
      {{ range.ImgNames }} 
        <img 
          loading="lazy" 
          class="img" 
          alt="food-{{.}}" 
          src="{{$.ImgUrl}}{{.}}" 
        /> 
      {{ end }}
      {{ range.ImgNames }} 
        <img 
          loading="lazy" 
          class="img" 
          alt="food-{{.}}" 
          src="{{$.ImgUrl}}{{.}}" 
        /> 
      {{ end }}
      {{ range.ImgNames }} 
        <img 
          loading="lazy" 
          class="img" 
          alt="food-{{.}}" 
          src="{{$.ImgUrl}}{{.}}" 
        /> 
      {{ end }}
      {{ range.ImgNames }} 
        <img 
          loading="lazy" 
          class="img" 
          alt="food-{{.}}" 
          src="{{$.ImgUrl}}{{.}}" 
        /> 
      {{ end }}
      {{ range.ImgNames }} 
        <img 
          loading="lazy" 
          class="img" 
          alt="food-{{.}}" 
          src="{{$.ImgUrl}}{{.}}" 
        /> 
      {{ end }}
      {{ range.ImgNames }} 
        <img 
          loading="lazy" 
          class="img" 
          alt="food-{{.}}" 
          src="{{$.ImgUrl}}{{.}}" 
        /> 
      {{ end }}
      {{ range.ImgNames }} 
        <img 
          loading="lazy" 
          class="img" 
          alt="food-{{.}}" 
          src="{{$.ImgUrl}}{{.}}" 
        /> 
      {{ end }}
      {{ range.ImgNames }} 
        <img 
          loading="lazy" 
          class="img" 
          alt="food-{{.}}" 
          src="{{$.ImgUrl}}{{.}}" 
        /> 
      {{ end }}
      {{ range.ImgNames }} 
        <img 
          loading="lazy" 
          class="img" 
          alt="food-{{.}}" 
          src="{{$.ImgUrl}}{{.}}" 
        /> 
      {{ end }}
      {{ range.ImgNames }} 
        <img 
          loading="lazy" 
          class="img" 
          alt="food-{{.}}" 
          src="{{$.ImgUrl}}{{.}}" 
        /> 
      {{ end }}
      {{ range.ImgNames }} 
        <img 
          loading="lazy" 
          class="img" 
          alt="food-{{.}}" 
          src="{{$.ImgUrl}}{{.}}" 
        /> 
      {{ end }}
      {{ range.ImgNames }} 
        <img 
          loading="lazy" 
          class="img" 
          alt="food-{{.}}" 
          src="{{$.ImgUrl}}{{.}}" 
        /> 
      {{ end }}
      {{ range.ImgNames }} 
        <img 
          loading="lazy" 
          class="img" 
          alt="food-{{.}}" 
          src="{{$.ImgUrl}}{{.}}" 
        /> 
      {{ end }}
    </main>
    <div class="fullscreen-overlay">
      <div class="fullscreen-image-container">
        <img src="" alt="full screen image" class="fullscreen-image" />
        <button class="close-button">&#10005;</button>
      </div>
    </div>
    <script type="module" src="./main.js"></script>
  </body>
</html>
