import Swiper from 'https://cdn.jsdelivr.net/npm/swiper@11/swiper-bundle.min.mjs'


const servicesSwiper = new Swiper('.services-swiper', {
    effect: 'coverflow',
    grabCursor: true,
    centeredSlides: true,
    loop: true,
    speed: 500,
    slidesPerView: 'auto',
    coverflowEffect: {
      rotate: 0,
      depth: 500,
      modifier: 0.5,
      slideShadows: false,
  },
  autoplay: {
    delay: 3000,
    pauseOnMouseEnter: true,
    disableOnInteraction: 0
},
    pagination: {
      el: '.swiper-pagination',
      clickable: true,
    },
    navigation: {
      nextEl: '.swiper-button-next',
      prevEl: '.swiper-button-prev',
    },
  });


  // setInterval(()=>{
  //   servicesSwiper.slideNext();
  // },10000);


  const customersSwiper = new Swiper('.customers-swiper', {
    effect: 'coverflow',
    grabCursor: true,
    centeredSlides: true,
    speed: 500,
    loop: true,
    spaceBetween: 40,
    slidesPerView: 'auto',
    coverflowEffect: {
      rotate: 0,
      depth: 0,
      modifier: 1,
      slideShadows: false,
    
  },
    pagination: {
      el: '.swiper-pagination',
      clickable: true,
    },
    navigation: {
      nextEl: '.swiper-button-next',
      prevEl: '.swiper-button-prev',
    },
  });

  const articlesSwiper = new Swiper('.articles-swiper', {
    effect: 'coverflow',
    grabCursor: true,
    centeredSlides: true,
    speed: 500,
    loop:true,
    spaceBetween: 30,
    slidesPerView: 'auto',
    coverflowEffect: {
      rotate: 0,
      depth: 0,
      modifier: 2,
      slideShadows: false,
    
  },

    pagination: {
      el: '.swiper-pagination',
      clickable: true,
    },
    navigation: {
      nextEl: '.swiper-button-next',
      prevEl: '.swiper-button-prev',
    },
  });