<template>
  <div class="bg-gradient-to-br from-gray-900 via-black to-gray-900 text-white min-h-screen">
    <!-- Hero Section -->
    <section
      id="home"
      class="relative min-h-screen flex items-center justify-center px-4 overflow-hidden"
    >
      <!-- Animated background gradient -->
      <div
        class="absolute inset-0 bg-gradient-to-r from-blue-900/30 via-purple-900/30 to-pink-900/30 animate-gradient-x"
      ></div>

      <!-- Floating particles effect -->
      <div class="absolute inset-0 overflow-hidden">
        <div class="particles"></div>
      </div>

      <div class="max-w-7xl mx-auto text-center relative z-20 py-20">
        <div class="space-y-8">
          <div class="relative inline-block">
            <h1
              ref="heroTitle"
              class="text-6xl md:text-8xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-blue-400 via-purple-400 to-pink-400 relative flex items-center justify-center gap-4"
            >
              Creative Developer
              <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                class="w-16 h-16 fill-current"
              >
                <path
                  d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm-5.5-2.5l7.51-3.49L17.5 6.5 9.99 9.99 6.5 17.5zm5.5-6.6c.61 0 1.1.49 1.1 1.1s-.49 1.1-1.1 1.1-1.1-.49-1.1-1.1.49-1.1 1.1-1.1z"
                  fill="url(#gradient)"
                />
                <defs>
                  <linearGradient id="gradient" x1="0%" y1="0%" x2="100%" y2="0%">
                    <stop offset="0%" style="stop-color: #60a5fa" />
                    <stop offset="50%" style="stop-color: #a78bfa" />
                    <stop offset="100%" style="stop-color: #f472b6" />
                  </linearGradient>
                </defs>
              </svg>
              <span
                class="absolute -bottom-2 left-0 w-full h-1 bg-gradient-to-r from-blue-400 via-purple-400 to-pink-400"
              ></span>
            </h1>
          </div>
          <p
            ref="heroSubtitle"
            class="text-xl md:text-2xl text-gray-300 max-w-3xl mx-auto leading-relaxed font-light"
          >
            Transforming ideas into elegant digital experiences through innovative design and
            cutting-edge technology.
          </p>
          <div ref="heroCTA" class="flex justify-center gap-6">
            <a
              href="#projects"
              class="px-8 py-4 bg-gradient-to-r from-blue-600 to-purple-600 rounded-full font-medium transition-all duration-300 transform hover:-translate-y-1 hover:shadow-xl hover:shadow-blue-500/25 cursor-pointer"
            >
              View Projects
            </a>
            <a
              href="#contact"
              class="px-8 py-4 border border-white/20 hover:border-white/40 rounded-full font-medium transition-all duration-300 transform hover:-translate-y-1 hover:bg-white/10 hover:shadow-xl cursor-pointer backdrop-blur-sm"
            >
              Contact Me
            </a>
          </div>
        </div>
      </div>
    </section>

    <!-- About Section -->
    <section
      id="about"
      ref="aboutSection"
      class="min-h-screen flex items-center justify-center px-4 py-20 relative overflow-hidden"
    >
      <!-- Background gradient -->
      <div
        class="absolute inset-0 bg-gradient-to-b from-transparent via-blue-900/10 to-transparent"
      ></div>

      <div class="max-w-7xl mx-auto relative z-10">
        <h2 ref="aboutTitle" class="text-5xl font-bold mb-16 text-center relative">
          About Me
          <span
            class="absolute -bottom-2 left-1/2 transform -translate-x-1/2 w-32 h-1 bg-gradient-to-r from-blue-400 via-purple-400 to-pink-400"
          ></span>
        </h2>
        <div ref="aboutContent" class="grid md:grid-cols-2 gap-16 items-center">
          <div class="space-y-8">
            <p class="text-lg text-gray-300 leading-relaxed">
              I'm a passionate web developer dedicated to crafting immersive digital experiences.
              With a blend of technical expertise and creative problem-solving, I transform complex
              challenges into intuitive, beautiful solutions that drive user engagement and business
              success.
            </p>
            <div class="grid md:grid-cols-2 gap-6">
              <!-- Frontend Card -->
              <div
                class="bg-white/5 p-6 rounded-xl border border-white/10 hover:border-blue-500/50 transition-all duration-300"
              >
                <div class="flex items-center mb-4">
                  <div class="bg-blue-500/20 p-3 rounded-lg mr-4">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-8 w-8 text-blue-400"
                      fill="none"
                      viewBox="0 0 24 24"
                      stroke="currentColor"
                    >
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
                      />
                    </svg>
                  </div>
                  <h3 class="text-xl font-semibold text-blue-400">Frontend</h3>
                </div>
                <div class="grid grid-cols-2 gap-4">
                  <!-- Vue.js -->
                  <div
                    class="flex flex-col items-center p-4 bg-blue-500/10 rounded-lg hover:bg-blue-500/20 transition-all h-[120px] justify-center"
                  >
                    <img
                      src="https://img.icons8.com/external-tal-revivo-color-tal-revivo/96/external-vuejs-an-open-source-javascript-framework-for-building-user-interfaces-and-single-page-applications-logo-color-tal-revivo.png"
                      alt="Vue.js"
                      class="h-12 w-12"
                    />
                    <span class="mt-2 text-sm text-blue-300">Vue.js</span>
                  </div>
                  <!-- React -->
                  <div
                    class="flex flex-col items-center p-4 bg-blue-500/10 rounded-lg hover:bg-blue-500/20 transition-all h-[120px] justify-center"
                  >
                    <img
                      src="https://img.icons8.com/plasticine/100/react.png"
                      alt="React"
                      class="h-12 w-12"
                    />
                    <span class="mt-2 text-sm text-blue-300">React</span>
                  </div>
                  <!-- Svelte -->
                  <div
                    class="flex flex-col items-center p-4 bg-blue-500/10 rounded-lg hover:bg-blue-500/20 transition-all h-[120px] justify-center"
                  >
                    <img
                      src="https://upload.wikimedia.org/wikipedia/commons/1/1b/Svelte_Logo.svg"
                      alt="Svelte"
                      class="h-12 w-12"
                    />
                    <span class="mt-2 text-sm text-blue-300">Svelte</span>
                  </div>
                  <!-- JavaScript -->
                  <div
                    class="flex flex-col items-center p-4 bg-blue-500/10 rounded-lg hover:bg-blue-500/20 transition-all h-[120px] justify-center"
                  >
                    <svg viewBox="0 0 128 128" class="h-12 w-12">
                      <path fill="#F0DB4F" d="M1.408 1.408h125.184v125.185H1.408z" />
                      <path
                        fill="#323330"
                        d="M116.347 96.736c-.917-5.711-4.641-10.508-15.672-14.981-3.832-1.761-8.104-3.022-9.377-5.926-.473-1.69-.526-2.643-.23-3.659.751-3.128 4.255-4.09 7.027-3.231 1.806.609 3.456 2.108 4.489 4.489 4.587-3.022 4.587-3.022 7.807-5.024-1.207-1.89-1.832-2.766-2.663-3.572-2.853-3.022-6.649-4.582-12.792-4.482l-3.188.405c-3.066.975-5.976 2.911-7.787 5.674-5.165 5.926-3.666 16.753 2.579 21.127 6.217 4.625 15.312 5.648 16.48 10.075 1.096 5.315-3.958 7.017-8.976 6.386-3.712-.792-5.139-2.688-7.141-6.133l-7.804 4.504c.902 2.053 1.942 2.938 3.496 4.705 7.524 7.596 26.432 7.229 29.84-4.333.119-.388.935-2.751.285-6.386zM69.462 58.943H57.753l-.048 30.272c0 6.438.333 12.34-.714 14.149-1.713 3.558-6.152 3.117-8.175 2.427-2.059-1.012-3.106-2.451-4.319-4.485-.333-.584-.583-1.036-.667-1.071l-9.52 5.83c1.583 3.249 3.915 6.069 6.902 7.901 4.462 2.678 10.459 3.499 16.731 2.059 4.082-1.189 7.604-3.652 9.448-7.401 2.666-4.915 2.094-10.864 2.07-17.444.06-10.735.001-21.468.001-32.237z"
                      />
                    </svg>
                    <span class="mt-2 text-sm text-blue-300">JavaScript</span>
                  </div>
                </div>
              </div>

              <!-- Backend Card -->
              <div
                class="bg-white/5 p-6 rounded-xl border border-white/10 hover:border-green-500/50 transition-all duration-300"
              >
                <div class="flex items-center mb-4">
                  <div class="bg-green-500/20 p-3 rounded-lg mr-4">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-8 w-8 text-green-400"
                      fill="none"
                      viewBox="0 0 24 24"
                      stroke="currentColor"
                    >
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M5 12h14M5 12a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v4a2 2 0 01-2 2M5 12a2 2 0 00-2 2v4a2 2 0 002 2h14a2 2 0 002-2v-4a2 2 0 00-2-2m-2-4h.01M17 16h.01"
                      />
                    </svg>
                  </div>
                  <h3 class="text-xl font-semibold text-green-400">Backend</h3>
                </div>
                <div class="grid grid-cols-2 gap-4">
                  <!-- Golang -->
                  <div
                    class="flex flex-col items-center p-4 bg-green-500/10 rounded-lg hover:bg-green-500/20 transition-all h-[120px] justify-center"
                  >
                    <img
                      src="https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg"
                      alt="Golang"
                      class="h-12 w-12"
                    />
                    <span class="mt-2 text-sm text-green-300">Golang</span>
                  </div>
                  <!-- Python -->
                  <div
                    class="flex flex-col items-center p-4 bg-green-500/10 rounded-lg hover:bg-green-500/20 transition-all h-[120px] justify-center"
                  >
                    <svg viewBox="0 0 128 128" class="h-12 w-12">
                      <path
                        fill="#FFD845"
                        d="M49.33 62h29.159c8.117 0 14.511-6.868 14.511-15.019v-27.798c0-7.912-6.632-13.856-14.555-15.176-5.014-.835-10.195-1.215-15.187-1.191-4.99.023-9.612.448-13.805 1.191-12.355 2.181-14.453 6.751-14.453 15.176v10.817h29v4h-40.224000000000004c-8.484 0-15.914 5.108-18.237 14.811-2.681 11.12-2.8 17.919 0 29.53 2.075 8.642 7.03 14.659 15.515 14.659h9.946v-13.048c0-9.637 8.428-17.952 18.33-17.952zm-1.838-39.11c-3.026 0-5.478-2.479-5.478-5.545 0-3.079 2.451-5.581 5.478-5.581 3.015 0 5.479 2.502 5.479 5.581-.001 3.066-2.465 5.545-5.479 5.545zM122.281 48.811c-2.098-8.448-6.103-14.811-14.599-14.811h-10.682v12.981c0 10.05-8.794 18.019-18.511 18.019h-29.159c-7.988 0-14.33 7.326-14.33 15.326v27.8c0 7.91 6.745 12.564 14.462 14.834 9.242 2.717 17.994 3.208 29.051 0 7.349-2.129 14.487-6.411 14.487-14.834v-11.126h-29v-4h43.682c8.484 0 11.647-5.776 14.599-14.66 3.047-9.145 2.916-17.799 0-29.529zm-41.955 55.606c3.027 0 5.479 2.479 5.479 5.547 0 3.076-2.451 5.579-5.479 5.579-3.015 0-5.478-2.502-5.478-5.579 0-3.068 2.463-5.547 5.478-5.547z"
                      />
                    </svg>
                    <span class="mt-2 text-sm text-green-300">Python</span>
                  </div>
                  <!-- Laravel -->
                  <div
                    class="flex flex-col items-center p-4 bg-green-500/10 rounded-lg hover:bg-green-500/20 transition-all h-[120px] justify-center"
                  >
                    <img
                      src="https://img.icons8.com/fluency/96/laravel.png"
                      alt="Laravel"
                      class="h-12 w-12"
                    />
                    <span class="mt-2 text-sm text-green-300">Laravel</span>
                  </div>
                  <!-- MySQL -->
                  <div
                    class="flex flex-col items-center p-4 bg-green-500/10 rounded-lg hover:bg-green-500/20 transition-all h-[120px] justify-center"
                  >
                    <svg viewBox="0 0 128 128" class="h-12 w-12">
                      <path
                        fill="#00618A"
                        d="M116.948 97.807c-6.863-.187-12.104.452-16.585 2.341-1.273.537-3.305.552-3.513 2.147.7.733.809 1.829 1.365 2.731 1.07 1.73 2.876 4.052 4.488 5.268 1.762 1.33 3.577 2.751 5.465 3.902 3.358 2.047 7.107 3.217 10.34 5.268 1.906 1.21 3.799 2.733 5.658 4.097.92.675 1.537 1.724 2.732 2.147v-.194c-.628-.8-.79-1.898-1.366-2.733l-2.537-2.537c-2.48-3.292-5.629-6.184-8.976-8.585-2.669-1.916-8.642-4.504-9.755-7.609l-.195-.195c1.892-.214 4.107-.898 5.854-1.367 2.934-.786 5.556-.583 8.585-1.365l4.097-1.171v-.78c-1.531-1.571-2.623-3.651-4.292-5.073-4.37-3.72-9.138-7.437-14.048-10.537-2.724-1.718-6.089-2.835-8.976-4.292-.971-.491-2.677-.746-3.318-1.562-1.517-1.932-2.342-4.382-3.511-6.633-2.449-4.717-4.854-9.868-7.024-14.831-1.48-3.384-2.447-6.72-4.293-9.756-8.86-14.567-18.396-23.358-33.169-32-3.144-1.838-6.929-2.563-10.929-3.513-2.145-.129-4.292-.26-6.438-.391-1.311-.546-2.673-2.149-3.902-2.927-4.894-3.092-17.448-9.817-21.072-.975-2.289 5.581 3.421 11.025 5.462 13.854 1.434 1.982 3.269 4.207 4.293 6.438.674 1.467.79 2.938 1.367 4.489 1.417 3.822 2.652 7.98 4.487 11.511.927 1.788 1.949 3.67 3.122 5.268.718.981 1.951 1.413 2.145 2.927-1.204 1.686-1.273 4.304-1.95 6.44-3.05 9.615-1.899 21.567 2.537 28.683 1.36 2.186 4.567 6.871 8.975 5.073 3.856-1.57 2.995-6.438 4.098-10.732.249-.973.096-1.689.585-2.341v.195l3.513 7.024c2.6 4.187 7.212 8.562 11.122 11.514 2.027 1.531 3.623 4.177 6.244 5.073v-.196h-.195c-.508-.791-1.303-1.119-1.951-1.755-1.527-1.497-3.225-3.358-4.487-5.073-3.556-4.827-6.698-10.11-9.561-15.609-1.368-2.627-2.557-5.523-3.709-8.196-.444-1.03-.438-2.589-1.364-3.122-1.263 1.958-3.122 3.542-4.098 5.854-1.561 3.696-1.762 8.204-2.341 12.878-.342.122-.19.038-.391.194-2.718-.655-3.672-3.452-4.683-5.853-2.554-6.07-3.029-15.842-.781-22.829.582-1.809 3.21-7.501 2.146-9.172-.508-1.666-2.184-2.63-3.121-3.903-1.161-1.574-2.319-3.646-3.124-5.464-2.09-4.731-3.066-10.044-5.267-14.828-1.053-2.287-2.832-4.602-4.293-6.634-1.617-2.253-3.429-3.912-4.683-6.635-.446-.968-1.051-2.518-.391-3.513.21-.671.508-1.951 1.171-2.341 1.132-.931 4.284.196 5.463.585 3.157 1.039 5.816 2.035 8.391 3.513 1.203.692 2.416 2.035 3.903 2.341h1.756c2.613.588 5.549.195 8.002.975 4.324 1.378 8.169 3.404 11.708 5.658 10.735 6.834 19.502 16.523 25.463 28.097 1.519 2.963 2.184 5.707 3.513 8.781 2.65 6.147 5.937 12.437 8.585 18.337 2.647 5.897 5.214 11.841 8.976 16.585 1.972 2.483 9.697 3.81 13.17 5.073 2.438.904 6.244 1.682 8.391 2.731 4.151 2.035 8.245 4.409 12.097 6.634 1.926 1.112 7.845 4.333 8.197 6.829z"
                      />
                    </svg>
                    <span class="mt-2 text-sm text-green-300">MySQL</span>
                  </div>
                </div>
              </div>

              <!-- Download Resume Card -->
              <div class="md:col-span-2 flex justify-center py-2">
                <a
                  href="/resume.pdf"
                  download
                  class="inline-flex items-center px-6 py-3 bg-gradient-to-r from-blue-600 to-purple-600 rounded-lg font-medium transition-all duration-300 transform hover:-translate-y-1 hover:shadow-xl hover:shadow-purple-500/25"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-5 w-5 mr-2"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
                    />
                  </svg>
                  Download Resume
                </a>
              </div>
            </div>
          </div>
          <div ref="profileImage" class="relative group">
            <div
              class="absolute -inset-2 bg-gradient-to-r from-blue-600 to-purple-600 rounded-xl blur-xl opacity-30 group-hover:opacity-50 transition-all duration-300"
            ></div>
            <img
              src="../assets/download.jpg"
              alt="Profile"
              class="relative rounded-xl shadow-2xl transform transition-transform duration-300 group-hover:scale-105"
            />
          </div>
        </div>
      </div>
    </section>

    <!-- Projects Section -->
    <section
      id="projects"
      ref="projectsSection"
      class="min-h-screen px-4 py-20 relative overflow-hidden"
    >
      <!-- Background gradient -->
      <div
        class="absolute inset-0 bg-gradient-to-b from-transparent via-blue-900/5 to-transparent"
      ></div>

      <div class="max-w-7xl mx-auto relative z-10">
        <h2 ref="projectsTitle" class="text-5xl font-bold mb-16 text-center relative">
          Featured Projects
          <span
            class="absolute -bottom-2 left-1/2 transform -translate-x-1/2 w-32 h-1 bg-gradient-to-r from-blue-400 via-purple-400 to-pink-400"
          ></span>
        </h2>
        <div ref="projectsGrid" class="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
          <div
            v-for="(project, index) in projects"
            :key="index"
            ref="projectCards"
            class="group relative overflow-hidden rounded-xl bg-white/5 border border-white/10 hover:border-blue-500/50 transition-all duration-500 hover:shadow-2xl hover:shadow-blue-500/20 hover:-translate-y-2"
          >
            <div
              class="absolute inset-0 bg-gradient-to-br from-blue-500/10 via-purple-500/10 to-pink-500/10 opacity-0 group-hover:opacity-100 transition-opacity duration-500"
            ></div>
            <div class="aspect-video overflow-hidden rounded-t-xl">
              <img
                :src="project.image"
                :alt="project.title"
                class="w-full h-full object-cover transform transition-all duration-500 group-hover:scale-110 group-hover:rotate-3"
              />
            </div>
            <div class="p-6 relative z-10">
              <h3
                class="text-xl font-semibold mb-2 text-blue-400 group-hover:text-blue-300 transition-colors duration-300"
              >
                {{ project.title }}
              </h3>
              <p
                class="text-gray-400 mb-4 group-hover:text-gray-300 transition-colors duration-300"
              >
                {{ project.description }}
              </p>
              <div class="flex flex-wrap gap-2">
                <span
                  v-for="(tech, techIndex) in project.technologies"
                  :key="techIndex"
                  class="px-3 py-1 text-sm bg-white/10 text-gray-300 rounded-full backdrop-blur-sm group-hover:bg-white/20 group-hover:text-white transition-all duration-300"
                >
                  {{ tech }}
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Contact Section -->
    <section
      id="contact"
      ref="contactSection"
      class="min-h-screen flex items-center justify-center px-4 py-20 relative"
    >
      <!-- Background gradient -->
      <div
        class="absolute inset-0 bg-gradient-to-t from-transparent via-purple-900/10 to-transparent"
      ></div>

      <div
        class="max-w-3xl mx-auto w-full backdrop-blur-lg bg-white/5 p-8 rounded-2xl border border-white/10"
      >
        <h2 ref="contactTitle" class="text-5xl font-bold mb-16 text-center relative">
          Get In Touch
          <span
            class="absolute -bottom-2 left-1/2 transform -translate-x-1/2 w-32 h-1 bg-gradient-to-r from-blue-400 via-purple-400 to-pink-400"
          ></span>
        </h2>
        <form ref="contactForm" class="space-y-8" @submit.prevent="handleSubmit">
          <div>
            <label class="block text-sm font-medium mb-2 text-gray-300">Name</label>
            <input
              v-model="formData.name"
              type="text"
              required
              class="w-full px-4 py-3 bg-white/5 border border-white/20 rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500 transition-all duration-300"
            />
          </div>
          <div>
            <label class="block text-sm font-medium mb-2 text-gray-300">Email</label>
            <input
              v-model="formData.email"
              type="email"
              required
              class="w-full px-4 py-3 bg-white/5 border border-white/20 rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500 transition-all duration-300"
            />
          </div>
          <div>
            <label class="block text-sm font-medium mb-2 text-gray-300">Message</label>
            <textarea
              v-model="formData.message"
              rows="4"
              required
              class="w-full px-4 py-3 bg-white/5 border border-white/20 rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500 transition-all duration-300"
            ></textarea>
          </div>
          <div
            v-if="submitStatus"
            :class="[
              'text-center mb-4 p-4 rounded-lg backdrop-blur-sm',
              submitStatus.success
                ? 'bg-green-500/10 text-green-400 border border-green-500/20'
                : 'bg-red-500/10 text-red-400 border border-red-500/20',
            ]"
          >
            {{ submitStatus.message }}
          </div>
          <button
            type="submit"
            :disabled="isSubmitting"
            class="w-full px-8 py-4 bg-gradient-to-r from-blue-600 to-purple-600 rounded-lg font-medium transition-all duration-300 transform hover:-translate-y-1 hover:shadow-xl hover:shadow-purple-500/25 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {{ isSubmitting ? 'Sending...' : 'Send Message' }}
          </button>
        </form>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
defineOptions({
  name: 'HomeView',
})

import { onMounted, ref, reactive } from 'vue'
import { gsap } from 'gsap'
import { ScrollTrigger } from 'gsap/ScrollTrigger'
import emailjs from '@emailjs/browser'

// Register ScrollTrigger plugin
gsap.registerPlugin(ScrollTrigger)

// Refs for animation targets
const heroTitle = ref<HTMLElement | null>(null)
const heroSubtitle = ref<HTMLElement | null>(null)
const heroCTA = ref<HTMLElement | null>(null)
const aboutSection = ref<HTMLElement | null>(null)
const aboutTitle = ref<HTMLElement | null>(null)
const aboutContent = ref<HTMLElement | null>(null)
const frontendSkills = ref<HTMLElement | null>(null)
const backendSkills = ref<HTMLElement | null>(null)
const profileImage = ref<HTMLElement | null>(null)
const projectsSection = ref<HTMLElement | null>(null)
const projectsTitle = ref<HTMLElement | null>(null)
const projectsGrid = ref<HTMLElement | null>(null)
const projectCards = ref<HTMLElement[]>([])
const contactSection = ref<HTMLElement | null>(null)
const contactTitle = ref<HTMLElement | null>(null)
const contactForm = ref<HTMLElement | null>(null)

const projects = [
  {
    title: 'Modern E-commerce Platform',
    description:
      'A full-featured online shopping platform with real-time inventory management and secure payment processing.',
    image:
      'https://images.unsplash.com/photo-1607799279861-4dd421887fb3?auto=format&fit=crop&q=80&w=1470',
    technologies: ['Vue.js', 'Tailwind', 'Node.js', 'MongoDB'],
  },
  {
    title: 'AI-Powered Analytics Dashboard',
    description:
      'Interactive data visualization platform featuring machine learning insights and predictive analytics.',
    image:
      'https://images.unsplash.com/photo-1551288049-bebda4e38f71?auto=format&fit=crop&q=80&w=1470',
    technologies: ['React', 'D3.js', 'Python', 'TensorFlow'],
  },
  {
    title: 'Social Media Management Suite',
    description:
      'Comprehensive tool for managing multiple social media accounts with scheduling and analytics capabilities.',
    image:
      'https://images.unsplash.com/photo-1460925895917-afdab827c52f?auto=format&fit=crop&q=80&w=1415',
    technologies: ['Next.js', 'TypeScript', 'GraphQL', 'AWS'],
  },
]

interface SubmitStatus {
  success: boolean
  message: string
}

const formData = reactive({
  name: '',
  email: '',
  message: '',
})

const isSubmitting = ref(false)
const submitStatus = ref<SubmitStatus | null>(null)

const handleSubmit = async () => {
  isSubmitting.value = true
  submitStatus.value = null

  try {
    const response = await emailjs.send(
      'service_joelyyy', // Replace with your EmailJS service ID
      'template_sucw2zq', // Replace with your EmailJS template ID
      {
        from_name: formData.name,
        from_email: formData.email,
        message: formData.message,
      },
      'UmWZjxSq8VSpPQiSX', // Replace with your EmailJS public key
    )

    if (response.status === 200) {
      submitStatus.value = {
        success: true,
        message: 'Message sent successfully! Thank you for reaching out.',
      }

      // Reset form
      formData.name = ''
      formData.email = ''
      formData.message = ''
    } else {
      throw new Error('Failed to send message')
    }
  } catch (err: unknown) {
    console.error('Error sending message:', err)
    submitStatus.value = {
      success: false,
      message: 'Failed to send message. Please try again later.',
    }
  } finally {
    isSubmitting.value = false
  }
}

// Initialize EmailJS
onMounted(() => {
  emailjs.init('UmWZjxSq8VSpPQiSX') // Updated with your actual public key

  // Hero Section Animations
  const heroTimeline = gsap.timeline()
  heroTimeline
    .from(heroTitle.value, {
      opacity: 0,
      y: 50,
      duration: 1,
      ease: 'power3.out',
    })
    .from(
      heroSubtitle.value,
      {
        opacity: 0,
        y: 30,
        duration: 1,
        ease: 'power3.out',
      },
      '-=0.5',
    )

  if (heroCTA.value?.children) {
    heroTimeline.from(
      heroCTA.value.children,
      {
        opacity: 0,
        y: 20,
        duration: 0.8,
        stagger: 0.2,
        ease: 'power3.out',
      },
      '-=0.5',
    )
  }

  // Scroll-triggered animations
  const sections = [
    {
      trigger: aboutSection.value,
      elements: [
        aboutTitle.value,
        aboutContent.value,
        frontendSkills.value,
        backendSkills.value,
        profileImage.value,
      ],
    },
    {
      trigger: projectsSection.value,
      elements: [projectsTitle.value, projectsGrid.value],
    },
    {
      trigger: contactSection.value,
      elements: [contactTitle.value, contactForm.value],
    },
  ]

  sections.forEach((section) => {
    gsap.from(section.elements, {
      scrollTrigger: {
        trigger: section.trigger,
        start: 'top 80%',
        toggleActions: 'play none none reverse',
      },
      opacity: 0,
      y: 50,
      stagger: 0.2,
      duration: 1,
      ease: 'power3.out',
    })
  })

  // Project cards animation
  projectCards.value.forEach((card) => {
    gsap.from(card, {
      scrollTrigger: {
        trigger: card,
        start: 'top bottom-=100',
        end: 'top center+=100',
        scrub: 0.5,
        toggleActions: 'play reverse play reverse',
      },
      opacity: 0,
      scale: 0.95,
      y: 50,
      rotateX: -5,
      transformOrigin: '50% 50%',
      duration: 0.6,
      ease: 'power1.out',
    })
  })
})
</script>

<style scoped>
.animate-gradient-x {
  background-size: 400% 400%;
  animation: gradient 15s ease infinite;
}

@keyframes gradient {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}

.particles {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: radial-gradient(circle, rgba(255, 255, 255, 0.1) 1px, transparent 1px);
  background-size: 50px 50px;
  animation: particleMove 60s linear infinite;
}

@keyframes particleMove {
  0% {
    transform: translateY(0);
  }
  100% {
    transform: translateY(-100%);
  }
}
</style>
