<template>
  <div class="daterangepicker-calendar">
    <div class="d-flex align-items-center">
      <div class="p-4" :class="calendarIndex == 1 ? '' : 'invisible'">
        <button type="button" class="btn btn-sm btn-light" @mousedown.prevent @click="goToPrevMonth">
          <font-awesome-icon icon="caret-left" fixed-width />
        </button>
      </div>
      <div class="p-1 col text-center">
        {{ displayMonth.format('MMMM YYYY') }}
      </div>
      <div class="p-1" :class="calendarIndex == calendarCount ? '' : 'invisible'">
        <button type="button" class="btn btn-sm btn-light" @mousedown.prevent @click="goToNextMonth">
          <font-awesome-icon icon="caret-right" fixed-width />
        </button>
      </div>
    </div>
    
    <div class="d-flex pl-2 justify-content-between text-center daterangepicker-calendar-row">
      <div v-for="day in daysOfFirstWeek" :key="day.format('D')" class="col-day">{{ day.format('ddd') }}</div>
    </div>

    <div class="d-flex pl-2 flex-wrap justify-content-between text-center daterangepicker-calendar-row">
      <div v-for="day in days" :key="day.format('M-D')" class="col-day" :class="dayClass(day)"
        @mouseover="dayMouseOver(day)" @mousedown.prevent @click="dayClick(day)">
        {{ day.format('D') }}
      </div>
    </div>

  </div>
</template>

<script>
import moment from 'moment'
import { library } from '@fortawesome/fontawesome-svg-core'
import { faCaretLeft, faCaretRight } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

library.add(faCaretLeft, faCaretRight)

export default {
  props: ['calendarIndex', 'calendarCount', 'month', 'startDate', 'endDate', 'step'],
  data: () => {
    return {}
  },
  computed: {
    displayMonth: function() {
      return moment.utc(this.month).add(this.calendarIndex - 1, 'month')
    },
    days: function() {
      let startDay = moment.utc(this.displayMonth).startOf('isoWeek')
      let endDay = moment.utc(this.displayMonth).endOf('month').endOf('isoWeek').startOf('day').add(7, 'day')
      let nDays = moment.duration(endDay.diff(startDay)).asDays()

      let days = []
      let day = 0
      while (day < nDays) {
        days.push(moment.utc(startDay).add(day, 'day'))
        day++
      }
      return days
    },
    daysOfFirstWeek: function() {
      return this.days.slice(0, 7)
    }
  },
  
  methods: {

    dayClass: function(day) {
      let classes = []

      // Hide days overflowing
      if (!day.isBetween(this.displayMonth, moment.utc(this.displayMonth).endOf('month'), 'days', '[]')) {
        classes.push('invisible')
      }
      // Class for days between startDate & endDate or is startDate (in case of startDate after endDate)
      if (day.isBetween(this.startDate, this.endDate, 'days', '[]') || day.isSame(this.startDate)) {
        classes.push('daterangepicker-range')
      }
      // Class for cursor if the step is selecting something
      if (this.step != null) {
        classes.push('daterangepicker-cursor-pointer')
      }

      return classes.join(' ')
    },
    dayMouseOver: function(day) {
      if (this.step != null) {
        this.selectDate(day)
      }
    },
    dayClick: function() {
      if (this.step != null) {
        this.nextStep()
      }
    },
    goToPrevMonth: function() {
      this.$emit('goToPrevMonth')
    },
    goToNextMonth: function() {
      this.$emit('goToNextMonth')
    },
    selectDate: function(date) {
      this.$emit('selectDate', date)
    },
    nextStep: function() {
      this.$emit('nextStep')
    }
  },
  watch: {},
  filters: {},
  components: { FontAwesomeIcon }
}
</script>

<style>
.daterangepicker-calendar-row {
  font-size: 12px;
  /*max-width: 230px;*/
}

.col-day {
  width: 14.28%;
  padding: 0.5rem 0;
  white-space: nowrap;
  cursor: default;
}

.daterangepicker-cursor-pointer {
  cursor: pointer;
}
</style>
