<template>
  <div class="container-fluid">
    <div class="animated fadeIn loading" v-if="isShowModel === false">
      <b-spinner style="width: 3rem; height: 3rem;" label="Large Spinner" type="grow"></b-spinner>
    </div>

    <div class="animated fadeIn font" v-if="isShowModel === true">
      <h1>HISTORY</h1>
      <hr class="my-4" />

      <b-card-group columns class="card-rows cols-2 mb-3">
        <b-card class="shadow mb-4 bg-white rounded">
          <BarColumn v-bind:model="TotalModel" />
        </b-card>
        <b-card class="shadow mb-4 bg-white rounded">
          <Pie v-bind:model="TotalModel" />
        </b-card>
      </b-card-group>

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
          placeholder="Search Sprint..."
          aria-label="Search"
          autocomplete="on"
        />
      </div>

      <b-card-group rows class="card-rows mb-3">
        <b-card class="shadow mb-4 bg-white rounded">
          <carousel
            :per-page="1"
            :scrollPerPage="true"
            :centerMode="true"
            :paginationEnabled="true"
            :navigationEnabled="true"
            :paginationPadding="5"
            class="mb-4"
          >
            <slide v-for="(model) in filteredSprintBurndownChart" :key="model.titleSprint">
              <burndownChart v-bind:model="model" />
            </slide>
          </carousel>
        </b-card>
      </b-card-group>

      <carousel
        :navigationEnabled="true"
        :perPageCustom="[[320, 1],[768,2]]"
        :centerMode="true"
        :scrollPerPage="false"
        :paginationEnabled="false"
      >
        <slide v-for="(model,index) in filteredSprintModel" :key="index">
          <div class="card cardsprit mr-1 ml-1 shadow">
            <div class="card-body">
              <div class="text-value">{{model.title}}</div>
              <p>{{model.startDate}} - {{ model.endDate}}</p>
              <button
                class="btn-hover color-8"
                @click="selectSprint(model)"
                v-b-modal.modal-xl
              >
                <i class="icon-chart font-2xl d-block"></i>
              </button>
            </div>
          </div>
        </slide>
      </carousel>

      <b-modal id="modal-xl" size="xl" title="Bootstrap-Vue" hide-footer hide-header centered>
        <Bar v-bind:model="select" />
      </b-modal>
    </div>
  </div>
</template>


<script>
import Bar from "@/components/history/Bar.vue";
import BarColumn from "@/components/history/BarColumn.vue";
import burndownChart from "@/components/burndownChart/burndownChart.vue";
import Pie from "@/components/history/Pie.vue";
import { mapGetters } from "vuex";
import { Carousel, Slide } from "vue-carousel";
import { BoardService } from "../../services/BoardService";
const boardService = new BoardService();
export default {
  data() {
    return {
      search: '',
      TotalModel: Object,
      burndown: Object,
      select: Object,
      SprintModel: {
        scoreOfSprint: Object
      },
      isShowModel: false
    };
  },
  mounted: function() {
    this.getHistory();
  },
  computed: {
    ...mapGetters({ token: "user/token", idBoard: "user/idBoard" }),
    filteredSprintModel() {
      let text = this.search.trim().toLowerCase()
      return this.SprintModel.scoreOfSprint.filter(index => {
        return index.title.toLowerCase().includes(text)
      });
    },

    filteredSprintBurndownChart: function() {
      let text = this.search.trim().toLowerCase()
      return this.burndown.filter(model => {
        return model.titleSprint.toLowerCase().includes(text)
      });
    }
  },
  components: {
    Bar,
    BarColumn,
    burndownChart,
    Pie,
    Carousel,
    Slide
  },
  methods: {
    selectSprint(model) {
      this.select = model;
    },
    getHistory() {
      boardService
        .fetchHistory({ idBoard: this.idBoard })
        .then(resp => {
          this.burndown = resp.data.burnDown.burnDownChart;
          this.TotalModel = resp.data.histories.ScoreTotal;
          this.SprintModel = {
            ...this.SprintModel,
            ...{
              scoreOfSprint: resp.data.histories.scoreOfSprint
            }
          };
          this.isShowModel = true;
        })
        .catch(err => {
          alert(err);
        });
    }
  }
};
</script>

<style >
.loading {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
.font h1 {
  font-size: xxx-large;
  margin-bottom: 10px;
}
.cardsprit {
  overflow: hidden;
  border-radius: 20px;
  background-color: #e3e8ed;
  height: 170px;
  text-shadow: 2px 2px 4px white;
  color: #3c4b64;
}

.btn-hover {
  width: 100px;
  font-weight: 600;
  color: #fff;
  margin: 5px;
  height: 50px;
  text-align: center;
  border: none;
  background-size: 300% 100%;
  border-radius: 50px;
  -o-transition: all 0.4s ease-in-out;
  -webkit-transition: all 0.4s ease-in-out;
  transition: all 0.4s ease-in-out;
}

.btn-hover:hover {
  background-position: 100% 0;
  -o-transition: all 0.4s ease-in-out;
  -webkit-transition: all 0.4s ease-in-out;
  transition: all 0.4s ease-in-out;
}

.btn-hover:focus {
  outline: none;
}
.btn-hover.color-8 {
  background-image: linear-gradient(
    to right,
    #29323c,
    #485563,
    #2b5876,
    #4e4376
  );
  box-shadow: 0 4px 15px 0 rgba(45, 54, 65, 0.75);
}
</style>

