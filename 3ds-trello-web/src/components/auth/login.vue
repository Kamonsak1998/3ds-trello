<template>
  <div class='login'>
    <div class='container'>
      <img class='bglogin' src='@/assets/circle.png' alt />
      <button class='buttons' @click='Auth()'>
        <span>Login</span>
      </button>
    </div>
  </div>
</template>

<script>
import { OAuth } from 'oauthio-web';
import { mapActions, mapGetters } from 'vuex';
export default {
  name: 'login',
  computed: {
    ...mapGetters({ token: 'user/token' })
  },
  methods: {
    ...mapActions({
      getToken: 'user/getToken',
      getUsername: 'user/getUsername',
      getIduser: 'user/getIduser',
      getAvatar: 'user/getAvatar'
    }),

    Auth() {
      const self = this;
      OAuth.initialize('HPINSnNzQSvEP7Bh5V0AVjTp6NI'); //www.oath.io ID : 3dsxtrello01@gmail.com || Password : 3dsinteractive ถ้าหากเกิดปัญหา oauth ลองเข้าไปเช็คที่นี่นะครับ
      var provider = 'trello';
      OAuth.popup(provider, { cache: true })
        .done(function(trello) {
          if (trello.provider == 'trello') {
            const token = trello.oauth_token;
            self.getToken(token);
            // self.$store.dispatch('token/getToken', token)
            trello
              .me()
              .done(function(response) {
                if (response.raw.avatarUrl !== null) {
                   self.getAvatar(response.raw.avatarUrl)
                }
                self.getUsername(response.name);
                self.getIduser(response.raw.id);
                if (self.token) {
                  self.$router.push('/dashboards');
                }
              })
              .fail(function(err) {
                alert(err);
              });
          }
        })
        .fail(function(err) {
          alert(err);
        });
    },
  }
};
</script>

<style lang='scss' >
.buttons {
  font-size: 18px;
  text-shadow: 2px 2px 4px #000000;
  line-height: 1;
  color: #fff;
  letter-spacing: 0.025em;
  background: linear-gradient(40deg, #ff9966, #ff6666, #cc66cc) !important;
  padding: 18px 0 11px;
  cursor: pointer;
  border: 0;
  border-radius: 2px;
  min-width: 120px;
  overflow: hidden;

  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.buttons span {
  display: block;
  position: relative;
  z-index: 10;
}

.buttons:after,
.buttons:before {
  padding: 36px 0 11px;
  content: '';
  position: absolute;
  top: 0;
  left: calc(-100% - 30px);
  height: calc(100% - 29px);
  width: calc(100% + 20px);
  color: #fff;
  border-radius: 2px;
  transform: skew(-25deg);
}

.buttons:after {
  background: #fff;
  transition: left 0.8s cubic-bezier(0.86, 0, 0.07, 1) 0.2s;
  z-index: 0;
  opacity: 0.8;
}

.buttons:before {
  background: linear-gradient(40deg, #ff9966, #ff6666, #cc66cc) !important;
  z-index: 5;
  transition: left 1s cubic-bezier(0.86, 0, 0.07, 1);
}

.buttons:hover:after {
  left: calc(0% - 10px);
  transition: left 0.8s cubic-bezier(0.86, 0, 0.07, 1);
}

.buttons:hover:before {
  left: calc(0% - 10px);
  transition: left 1s cubic-bezier(0.86, 0, 0.07, 1);
}
.bglogin {
  width: 620px;
  height: 620px;
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
</style>
