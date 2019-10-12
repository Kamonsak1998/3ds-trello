<template>
  <div class="container pt-5">
    <div class="animated fadeIn loading" v-if="isShowModel === false">
      <b-spinner style="width: 3rem; height: 3rem;" label="Large Spinner" type="grow"></b-spinner>
    </div>
    
    <div class="row" v-if="isShowModel === true">
      <div class="input-group input-group-lg my-3">
        <div class="input-group-prepend">
          <span class="input-group-text">
            <i class="icon-magnifier"></i>
          </span>
        </div>
        <input
          type="text"
          id="search"
          class="form-control"
          v-model="search"
          placeholder="Search Board..."
          aria-label="Search"
          autocomplete="on"
        />
      </div>
      <div v-for="(result,index) in filteredBoardModel" :key="index" class="col-sm-4">
        <b-card
          overlay
          :img-src="result.imageBackground"
          img-alt="Card Image"
          text-variant="white"
          :title="result.boardName"
          style="max-width: 30rem;"
          align="center"
          class="imgbg shadow-lg block"
          @click="setboard(results,index)"
        >
          <h3 class="animate-text text-animate">
            <b-card-text>Choose this project</b-card-text>
          </h3>
        </b-card>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions, mapGetters } from "vuex";
import { BoardService } from "../../services/BoardService";

const boardService = new BoardService();
export default {
  mounted: function() {
    this.getBoardtrello();
  },
  computed: {
    ...mapGetters({ token: "user/token" , idBoard: "user/idBoard" }),
    filteredBoardModel() {
      let text = this.search.trim().toLowerCase()
      return this.results.filter(result => {
        return result.boardName.toLowerCase().includes(text)
      });
    },
  },
  data() {
    return {
      search:'',
      results: [],
      resultss: [],
      isShowModel: false
    };
  },

  methods: {
    search_text() {
      var inside = this;
      this.results = this.resultss.filter(function(results) {
        if (
          results.boardName
            .toLowerCase()
            .indexOf(inside.search.text.toLowerCase()) != "-1"
        ) {
          return results;
        }
      });
    },
    ...mapActions({
      getBoard: "user/getBoard",
      getNameBoard: "user/getNameBoard"
    }),
    
    setboard(result, index) {
      const boardid = result[index].idBoard;
      const nameBoard = result[index].boardName;
      this.getBoard(boardid);
      this.getNameBoard(nameBoard);
      this.$router.push("/feature");
    },
    getBoardtrello() {
      var inside = this;
      boardService
        .fetchDashboard()
        .then(Response => {
          inside.resultss = Response.data;
          inside.results = Response.data;
          this.isShowModel = true;
        })
        .catch(err => {
          alert(err);
        });
    }
  }
};
</script>

<style lang="scss">
.loading {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
.imgbg {
  border: none;
  width: 100%;
  height: 180px;
  border-radius: 25px;
  overflow: hidden;
  cursor: pointer;
  color: white;
  text-shadow: 2px 2px 4px #000000;
  background: linear-gradient(40deg, #ff9966, #ff6666, #cc66cc) !important;
}
.card-img {
  border-radius: 25px;
  width: 100%;
  height: 100%;
}
.text-animate {
  font-weight: 100;
  font-style: italic;
  transform: translateX(200px);
}
.animate-text {
  opacity: 0;
  transition: all 0.6s ease-in-out;
}
.block:hover .animate-text {
  transform: translateX(0);
  opacity: 1;
}
.block:hover {
  z-index: 100;
  -webkit-animation: scale 0.3s linear;
  -moz-animation: scale 0.3s linear;
  animation: scale 0.3s linear;
  transform-origin: 50% 50%;
  animation-fill-mode: forwards;
}
@keyframes scale {
  0% {
    transform: scale(1);
  }
  100% {
    transform: scale(1.1);
    -webkit-box-shadow: 10px 10px 60px 10px rgba(0, 0, 0, 0.1);
    -moz-box-shadow: 10px 10px 60px 10px rgba(0, 0, 0, 0.1);
    box-shadow: 10px 10px 60px 10px rgba(0, 0, 0, 0.1);
  }
}

@keyframes scaledown {
  0% {
    transform: scale(1.1);
  }
  100% {
    transform: scale(1);
  }
}
</style>
