<template>
    <div class="video-container position-relative mx-auto">
        <video
          ref="videoPlayer" 
          :src="src" 
          playsinline 
          loop
          autoplay
          @canplay="onVideoCanPlay"
          @pause="onPause"
          @play="onPlay"
        ></video>
        <div class="controls position-absolute w-100 h-100 d-flex justify-content-center align-items" @click="togglePlay">
            <div class="volume-control">
                    <svg width="24" data-e2e="" height="24" viewBox="0 0 48 48" fill="#fff" xmlns="http://www.w3.org/2000/svg">
                        <path fill-rule="evenodd" clip-rule="evenodd" d="M20.3359 8.37236C22.3296 7.04325 25 8.47242 25 10.8685V37.1315C25 39.5276 22.3296 40.9567 20.3359 39.6276L10.3944 33H6C4.34314 33 3 31.6568 3 30V18C3 16.3431 4.34315 15 6 15H10.3944L20.3359 8.37236ZM21 12.737L12.1094 18.6641C11.7809 18.8831 11.3948 19 11 19H7V29H11C11.3948 29 11.7809 29.1169 12.1094 29.3359L21 35.263V12.737ZM32.9998 24C32.9998 21.5583 32.0293 19.3445 30.4479 17.7211C30.0625 17.3255 29.9964 16.6989 30.3472 16.2724L31.6177 14.7277C31.9685 14.3011 32.6017 14.2371 33.0001 14.6195C35.4628 16.9832 36.9998 20.3128 36.9998 24C36.9998 27.6872 35.4628 31.0168 33.0001 33.3805C32.6017 33.7629 31.9685 33.6989 31.6177 33.2724L30.3472 31.7277C29.9964 31.3011 30.0625 30.6745 30.4479 30.2789C32.0293 28.6556 32.9998 26.4418 32.9998 24ZM37.0144 11.05C36.6563 11.4705 36.7094 12.0995 37.1069 12.4829C40.1263 15.3951 42.0002 19.4778 42.0002 23.9999C42.0002 28.522 40.1263 32.6047 37.1069 35.5169C36.7094 35.9003 36.6563 36.5293 37.0144 36.9498L38.3109 38.4727C38.6689 38.8932 39.302 38.9456 39.7041 38.5671C43.5774 34.9219 46.0002 29.7429 46.0002 23.9999C46.0002 18.2569 43.5774 13.078 39.7041 9.43271C39.302 9.05421 38.6689 9.10664 38.3109 9.52716L37.0144 11.05Z"></path>
                    </svg>
                <input type="range" min="0" max="1" step="0.01" v-model="volume" @input="changeVolume" />
            </div>
            <button class="play" :class="isPlaying ? 'animate' : ''">
                <font-awesome-icon :icon="isPlaying ? 'pause' : 'play'" />
            </button>
        </div>
        <section class="section-wrapper position-absolute">
            <div>
                <div class="wrapper-svg">
                    <svg class="style-svg" viewBox="0 0 48 48" fill="currentColor" xmlns="http://www.w3.org/2000/svg" width="2em" height="2em">
                        <path d="M24 9.44c3.2-4.03 7.61-5.56 12-4.67 2.31.46 5.59 2.28 7.75 5.48 2.26 3.32 3.21 7.99.98 13.85-1.75 4.57-5.5 8.83-9.28 12.2a56.6 56.6 0 0 1-10.52 7.47l-.93.49-.93-.49a56.6 56.6 0 0 1-10.52-7.47c-3.78-3.37-7.53-7.63-9.28-12.2-2.23-5.86-1.28-10.53.98-13.85C6.4 7.05 9.69 5.23 12 4.77c4.39-.9 8.8.64 12 4.67Z" fill="#fff" fill-opacity="0.9" shape-rendering="crispEdges"></path>
                    </svg>
                </div>
                <span class="info-text d-block">562K</span>
            </div>
            <div>
                <div class="wrapper-svg">
                    <svg class="style-svg" viewBox="0 0 48 48" fill="currentColor" xmlns="http://www.w3.org/2000/svg" width="2em" height="2em">
                        <path fill-rule="evenodd" clip-rule="evenodd" d="M38.5 35.31c4.1-4.11 6.5-8.4 6.5-13.38C45 11.8 35.73 3.6 24.3 3.6S3.6 11.8 3.6 21.93C3.6 32.05 13.17 39 24.6 39v3.36c0 1.06 1.1 1.74 2.04 1.24 2.92-1.59 8.33-4.76 11.85-8.29ZM14.23 19.46a2.95 2.95 0 0 1 2.96 2.93 2.95 2.95 0 0 1-2.96 2.94 2.95 2.95 0 0 1-2.95-2.94 2.95 2.95 0 0 1 2.95-2.93Zm13.02 2.93a2.95 2.95 0 0 0-2.96-2.93 2.95 2.95 0 0 0-2.96 2.93 2.95 2.95 0 0 0 2.96 2.94 2.95 2.95 0 0 0 2.96-2.94Zm7.1-2.93a2.95 2.95 0 0 1 2.95 2.93 2.95 2.95 0 0 1-2.96 2.94 2.95 2.95 0 0 1-2.95-2.94 2.95 2.95 0 0 1 2.95-2.93Z" fill="#fff" fill-opacity="0.9"></path>
                    </svg>
                </div>
                <span class="info-text d-block">3465</span>
            </div>
            <div>
                <div class="wrapper-svg">
                    <svg class="style-svg" viewBox="0 0 48 48" fill="currentColor" xmlns="http://www.w3.org/2000/svg" width="2em" height="2em">
                        <path fill-rule="evenodd" clip-rule="evenodd" d="M25.56 4.07a1.98 1.98 0 0 0-2.15-.42 1.95 1.95 0 0 0-1.21 1.8v8.34c-5.4.35-10.04 2.2-13.43 5.68C4.97 23.35 3 29.03 3 36.19c0 .79.48 1.5 1.22 1.8.73.3 1.58.13 2.14-.42 3.34-3.31 7.65-4.56 11.25-4.95 1.8-.2 3.37-.18 4.5-.1h.09v9.03c0 .78.46 1.48 1.18 1.79.72.3 1.56.16 2.13-.37l18.87-17.49a1.94 1.94 0 0 0 .04-2.8L25.56 4.07Z" fill="#fff" fill-opacity="0.9" shape-rendering="crispEdges"></path>
                    </svg>
                </div>
            <span class="info-text d-block">18.9K</span>
            </div>
        </section>      
    </div>
</template>

<script>
export default {
  props: {
    src: {
      type: String,
      required: true,
    },
  },
  data() {
    return {
      observer: null,
      isUserInteracted: false, // Track if the user has interacted
      isPlaying: false,
      volume: 0.5, // Default volume
    };
  },
  mounted() {
    this.initIntersectionObserver();
    this.autoplayVideo(); // Attempt to autoplay on mount
  },
  beforeUnmount() {
    if (this.observer) {
      this.observer.disconnect(); // Clean up observer on component destroy
    }
  },
  methods: {
    initIntersectionObserver() {
      this.observer = new IntersectionObserver(entries => {
        entries.forEach(entry => {
          if (entry.isIntersecting) {
            // Play the video if the user has interacted
            if (this.isUserInteracted) {
              this.playVideo();
            }
          } else {
            this.pauseVideo();
          }
        });
      });
      this.observer.observe(this.$refs.videoPlayer);
    },
    autoplayVideo() {
      this.playVideo(); // Start playing muted video
    },
    playVideo() {
      this.$refs.videoPlayer.play().then(() => {
        this.isPlaying = true;
      }).catch(error => {
        console.error('Error trying to play video:', error);
      });
    },
    pauseVideo() {
      this.$refs.videoPlayer.pause();
      this.isPlaying = false;
    },
    togglePlay() {
      if (this.isPlaying) {
        this.pauseVideo();
      } else {
        this.playVideo();
      }
    },
    changeVolume() {
      this.$refs.videoPlayer.volume = this.volume;
    },
    onVideoCanPlay() {
      console.log('Video is ready to play.');
    },
    onPause() {
      console.log('Video is paused.');
    },
    onPlay() {
      console.log('Video is playing.');
    },
  },
};
</script>
