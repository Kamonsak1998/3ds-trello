<template>
  <b-container class="bv-example-row row-setdate">
    <div class="animated fadeIn loading" v-if="isShowModel === false">
      <b-spinner style="width: 3rem; height: 3rem;" label="Large Spinner" type="grow"></b-spinner>
    </div>
    <div class="card card-setdate" v-if="isShowModel === true">
      <b-alert
        variant="success"
        dismissible
        :show="dismissCountDown"
        @dismissed="dismissCountDown=0"
        @dismiss-count-down="countDownChanged"
      >Save Success</b-alert>
      <b-row>
        <b-col v-for="calendarIndex in calendarCount" :key="calendarIndex">
          <date-range-picker-calendar
            :calendarIndex="calendarIndex"
            :calendarCount="calendarCount"
            :month="month"
            :startDate="startDate"
            :endDate="endDate"
            :step="step"
            v-on:goToPrevMonth="goToPrevMonth"
            v-on:goToNextMonth="goToNextMonth"
            v-on:selectDate="selectDate"
            v-on:nextStep="nextStep"
          />
        </b-col>
        <b-col class="col-setdate">
          <h2 class="pb-4">Set Date Time Stamp</h2>
          <p>Start Sprint</p>
          <div class="form-group form-inline flex-nowrap">
            <input
              type="text"
              class="form-control w-100 daterangepicker-date-input"
              ref="startDate"
              :value="startDate | dateFormat"
              :disabled="validated"
              @click="reset"
              @focus="step = 'selectStartDate'"
              @blur="inputDate"
            />
          </div>
          <br />
          <p>Sprint Period (Day)</p>
          <input
            name="total"
            type="text"
            class="form-control w-100 daterangepicker-date-input"
            pattern="^[1-9]+"
            ref="endDate"
            placeholder="1 - 30"
            :disabled="validated"
            v-model="total"
            v-validate="'required|numeric|max:2'"
            :class="{ 'is-invalid': submitted && errors.has('total') }"
          />
          <br />
          <div
            v-if="submitted && errors.has('total')"
            class="invalid-feedback"
          >{{ errors.first('total') }}</div>
          <br />
          <div class="form-group form-inline justify-content-end mb-0">
            <button type="button" class="btn btn-light mr-2" @click="reset">Reset</button>
            <button
              type="button"
              class="btn btn-primary submitbtn"
              @click="submit"
              :disabled="validated"
            >Submit</button>
          </div>
        </b-col>
      </b-row>
    </div>
    <b-card-group columns class="card-rows cols-2 mb-3" v-if="isShowModel === true">
      <setscore
        :point="pointt"
        @input="(newpoint) => {pointt = newpoint}"
        :model="points"
        class="scoreCard"
        ref="score"
      ></setscore>
      <selectlist
        :model="lists"
        :selectListed="selectListed"
        class="listCard"
        :listed="listed"
        ref="select"
      >{{selectListed}}</selectlist>
    </b-card-group>
  </b-container>
</template>

<script>
import setscore from "@/components/Setting/setscore";
import selectlist from "@/components/Setting/selectlist";
import moment from "moment";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faCaretRight } from "@fortawesome/free-solid-svg-icons";
import DateRangePickerCalendar from "./DateRangePickerCalendar";
import { mapGetters } from "vuex";
import { BoardService } from "../../services/BoardService";

const boardservice = new BoardService();

library.add(faCaretRight);

export default {
  provide() {
    return { parentValidator: this.$validator };
  },
  props: {
    calendarCount: {
      type: Number,
      default: 2
    },

    ranges: {
      type: Object,
      default: function() {
        return {};
      }
    }
  },
  data() {
    return {
      dismissSecs: 5,
      dismissCountDown: 0,
      pointt: [],
      selectListed: [],
      listed: Object,
      isShowModel: false,
      validated: false,
      total: "",
      startDate: moment.utc(),
      endDate: moment.utc(),
      rangeSelect: null,
      month: moment
        .utc()
        .subtract(1, "month")
        .startOf("month"),
      step: null,
      submitted: false
    };
  },
  mounted: function() {
    this.checkDate();
  },

  computed: {
    ...mapGetters({
      idBoard: "user/idBoard",
      nameBoard: "user/nameBoard",
      token: "user/token"
    }),
    nextMonth: function() {
      return moment.utc(this.month).add(1, "month");
    },
    // For multi prop watchers
    range: function() {
      return this.startDate, this.endDate;
    }
  },

  methods: {
    focusInput() {
      setTimeout(() => {
        this.$refs.startDate.focus();
      }, 200);
    },
    checkDate: function() {
      boardservice
        .fetchchecksetting({ idBoard: this.idBoard })
        .then(res => {
          this.isShowModel = true;
          this.points = res.data.scoreSize;
          this.lists = res.data.lists.list;
          this.listed = res.data.lists.selectList;
          if (res.data.date.status == true) {
            this.startDated = moment.utc(res.data.date.startDate, "YYYY/MM/DD");
            this.startDate = this.startDated;
            this.validated = res.data.date.status;
            // this.totaled = res.data.date.sprintDay;
            this.total = res.data.date.sprintDay;
          }
        })
        .catch(err => {
          alert(err);
        });
    },
    clear: function() {
      this.total = "";
    },

    reset: function() {
      this.validated = false;
      this.startDate = moment.utc();
      this.endDate = moment.utc();
      this.total = "";
      this.$refs.score.reset();
      this.focusInput();
    },

    goToPrevMonth: function() {
      this.month = moment.utc(this.month).subtract(1, "month");
    },
    goToNextMonth: function() {
      this.month = moment.utc(this.month).add(1, "month");
    },
    selectDate: function(date) {
      if (this.step === "selectStartDate") {
        this.startDate = date;
      } else if (this.step === "endDate") {
        this.endDate = date;
      }
    },
    // Step flow for date range selections
    nextStep: function() {
      if (this.step === "selectStartDate") {
        this.step = "selectEndDate";
        this.$refs.endDate.focus();
      } else if (this.step === "endDate") {
        this.step = null;
        this.$refs.endDate.blur();
      }
    },

    // Try to update the step date from an input value
    inputDate: function(input) {
      let date = moment.utc(input.target.value, "YYYY/MM/DD");
      if (date.isValid()) {
        this.selectDate(date);
      }
      this.nextStep();
    },
    formValidate() {
      return this.$refs.select.formValidate(), this.$refs.score.formValidate();
    },
    countDownChanged(dismissCountDown) {
        this.dismissCountDown = dismissCountDown
      },
    submit: function() {
      this.submitted = true;
      this.$validator.validate().then(valid => {
        this.formValidate();
        this.totaled = parseInt(this.total);
        if (valid) {
          this.dismissCountDown = this.dismissSecs
          boardservice
            .fetchsettingdata({
              sprintDate: {
                startDate: this.startDate,
                sprintDay: this.totaled,
                endDate: this.endDate,
                idBoard: this.idBoard
              },
              scoreSize: {
                Points: [
                  parseFloat(this.pointt[0]),
                  parseFloat(this.pointt[1]),
                  parseFloat(this.pointt[2]),
                  parseFloat(this.pointt[3]),
                  parseFloat(this.pointt[4]),
                  parseFloat(this.pointt[5]),
                  parseFloat(this.pointt[6]),
                  parseFloat(this.pointt[7])
                ]
              },
              selectList: this.selectListed
            })
            .then(() => {
              // this.$router.push("/feature");
            })
            .catch(err => {
              if (err) {
                alert("Sorry Connection not found");
              }
            });
        }
      });
    }
  },
  watch: {
    total: function(value) {
      this.endDate = moment
        .utc(this.startDate, "YYYY/MM/DD")
        .add(1 * value - 1, "days");
    }
  },
  filters: {
    dateFormat: function(value) {
      return value ? value.format("YYYY/MM/DD") : "";
    }
  },
  components: { DateRangePickerCalendar, setscore, selectlist }
};
</script>

<style>
/* Custom row */
.daterangepicker-row {
  margin: -0.5rem;
}
.daterangepicker-date-input {
  min-width: 120px;
}

/* Select menus border */
.daterangepicker-range-border {
  border-color: #17a2b8 !important;
}

.daterangepicker-range {
  background-color: red !important;
  color: #ffffff;
}

/* Date input focus */
.daterangepicker-date-input:focus {
  border-color: #17a2b8 !important;
  box-shadow: 0 0 0 0.2rem rgba(23, 162, 184, 0.25) !important;
}

.col-setdate {
  padding: 40px;
}

.row-setdate {
  display: -ms-flexbox;
  display: flex;
  -ms-flex-wrap: wrap;
  flex-wrap: wrap;
}

.card-setdate {
  border-radius: 10px;
}

.loading {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
</style>
