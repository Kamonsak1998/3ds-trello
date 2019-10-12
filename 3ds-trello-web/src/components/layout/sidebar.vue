<template>
  <div class="sidebar text-left">
    <nav class="sidebar-nav">  
      <ul class="nav">
        <li class="nav-item">
          <router-link class="nav-link" :to="{name : 'dashboards'}">
            <i class="nav-icon icon-home"></i> Dashboard
          </router-link>
        </li>
        <li class="nav-item"></li>
        <li class="nav-item nav-dropdown" v-if="idBoard">
          <router-link class="nav-link nav-dropdown-toggle" :to="{name : 'feature'}">
            <i class="nav-icon icon-list"></i> Feature
          </router-link>
          <ul class="nav-dropdown-items">
            <li class="nav-item ">
              <router-link class="nav-link" :to="{name : 'charts'}">
                <i class="nav-icon icon-chart"></i> History
              </router-link>
            </li>
            <li class="nav-item">
              <router-link class="nav-link" :to="{name : 'leaderboard'}">
                <i class="nav-icon icon-book-open"></i> LeaderBoard
              </router-link>
            </li>
            <li class="nav-item">
              <router-link class="nav-link" :to="{name : 'setdatetime'}">
                <i class="nav-icon icon-settings"></i> Setting
              </router-link>
            </li>
          </ul>
        </li>
        <li class="nav-item nav-dropdown d-lg-none">
          <a class="nav-link nav-dropdown-toggle" href="#">
            <i class="nav-icon icon-user"></i> User
          </a>
          <ul class="nav-dropdown-items">
            <li class="nav-item" v-if="token">
              <a class="nav-link" href @click=" logout">
                <i class="nav-icon icon-logout"></i>Logout
              </a>
            </li>
            <li class="nav-item" v-else>
              <router-link class="nav-link" :to="{name : 'login'}">
                <i class="nav-icon icon-login"></i> Login
              </router-link>
            </li>
          </ul>
        </li>
      </ul>
    </nav>
    <button class="sidebar-minimizer brand-minimizer" type="button"></button>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import { OAuth } from "oauthio-web";
export default {
  computed: {
    ...mapGetters({token: 'user/token' , idBoard:'user/idBoard'}),
  },
  methods: {
      logout: function() {
      OAuth.clearCache();
      this.$store.dispatch('user/logout');
    }
  }
};
</script>

<style>
</style>
