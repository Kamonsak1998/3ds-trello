<template>
  <div id="app">
    <Header />
    <div class="app-body">
      <Sidebar />
      <main class="main p-4">
        <div class="container-fluid">
          <div class="animated fadeIn">
            <router-view />
          </div>
        </div>
      </main>
    </div>
    <Footer />
  </div>
</template>

<script>
import Header from "./components/layout/header.vue";
import Sidebar from "./components/layout/sidebar.vue";
import Footer from "./components/layout/footer.vue";
export default {
  name: "app",
  components: {
    Sidebar,
    Header,
    Footer
  }
};
</script>

<style>
#app {
  background-image: url("./assets/Bgver2.png");
  background-position: center center;
  background-repeat: no-repeat;
  background-size: cover;
  background-attachment: fixed;
  height: 100%;
  width: 100%;
  display: flex;
  min-height: 100vh;
  flex-direction: column;
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

@media only screen and (max-width: 768px) {
  #app {
    background-image: url("./assets/BgMobile.png");
    background-position: center center;
    background-repeat: no-repeat;
    background-size: cover;
    background-attachment: fixed;
    height: 100%;
    width: 100%;
  }
  .main .container-fluid {
    padding: 0 0px;
  }
}
.main {
  flex: 1;
}
@font-face {
  font-family: kano;
  src: url("./assets/montserrat/Montserrat-Regular.ttf");
  /* font-weight: bold; */
}
* {
  font-family: kano;
}
</style>
