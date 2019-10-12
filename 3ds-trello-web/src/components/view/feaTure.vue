<template>
  <div class="container pt-5">
    <div class="animated fadeIn loading" v-if="isShowModel === false">
      <b-spinner style="width: 3rem; height: 3rem;" label="Large Spinner" type="grow"></b-spinner>
    </div>
    <div v-if="isShowModel === true">
      <div class="row">

        <div class="col-sm-4">
          <router-link style="text-decoration:none" :to="{name : 'charts'}">
            <div class="card text-white bg-info">
              <div class="card-body card-body-feature History-f pb-0">
                <div class="card-body-feature">
                  <i class="icon-chart font-2xl d-block"></i>
                  <div class="text-value">History</div>
                </div>
              </div>
            </div>
          </router-link>
        </div>

         <div class="col-sm-4">
          <router-link style="text-decoration:none" :to="{name : 'leaderboard'}">
            <div class="card text-white bg-primary">
              <div class="card-body card-body-feature leaderboard-f pb-0">
                <div class="card-body-feature">
                  <i class="icon-book-open font-2xl d-block"></i>
                  <div class="text-value">Leader Board</div>
                </div>
              </div>
            </div>
          </router-link>
        </div>

        <div class="col-sm-4">
          <router-link style="text-decoration:none" :to="{name : 'setdatetime'}">
            <div class="card text-white bg-warning">
              <div class="card-body card-body-feature Set-date-f pb-0">
                <div class="card-body-feature">
                  <i class="icon-settings font-2xl d-block"></i>
                  <div class="text-value">Setting</div>
                </div>
              </div>
            </div>
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import { BoardService } from "../../services/BoardService";
const boardService = new BoardService();

export default {
  data() {
    return {
      isShowModel: false
    };
  },
  mounted: function() {
    this.checkidBoard();
  },
  computed: {
    ...mapGetters({ token: "user/token", idBoard: "user/idBoard" })
  },
  methods: {
    checkidBoard() {
      this.checkDate();
      this.setmember();
    },
    checkDate: function() {
      boardService
        .fetchchecksetting({ idBoard: this.idBoard })
        .then(res => {
          if (res.data.date.status == false) {
            this.$router.push("/setdatetime");
          } else {
            this.isShowModel = true;
          }
        })
        .catch(err => {
          alert(err);
        });
    },
    setmember() {
      boardService
        .fetchsetmember({
          idBoard: this.idBoard
        })
        .catch(err => {
          alert(err);
        });
    }
  }
};
</script>

<style>
.loading {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
.card.bg-primary {
  border-color: #187da0;
}

.bg-primary:hover {
  box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
}
/***********/
.card.bg-info {
  border-color: #2eadd3;
}
.bg-info:hover {
  box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
}
/**********/
.card.bg-warning {
  border-color: #c69500;
}
.bg-warning:hover {
  box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
}
/***********/
.card-body-feature {
  width: 100%;
  height: 120px;
  padding-top: 17px;
}
.text-value {
  font-size: 20px;
  font-weight: 600;
}
.boardff {
  color: white;
  text-align: center;
  text-decoration: none;
  display: inline-block;
}
</style>
