<template>
    <div class="animated fadeIn">
        <div>
            <p>
                <label>
                    <b>Select Theme</b>
                </label> &nbsp;
                <select @change="updateTheme">
                    <option value="palette1">Theme 1</option>
                    <option value="palette2">Theme 2</option>
                    <option value="palette3">Theme 3</option>
                    <option value="palette4">Theme 4</option>
                    <option value="palette5">Theme 5</option>
                    <option value="palette6">Theme 6</option>
                    <option value="palette7">Theme 7</option>
                    <option value="palette8">Theme 8</option>
                    <option value="palette9">Theme 9</option>
                    <option value="palette10">Theme 10</option>
                </select>
            </p>
            <apexchart height="350" type="bar" :options="chartOptions" :series="series"></apexchart>
        </div>

    </div>
</template>

<script>
    export default {
        name: "Bar",
        props: {
            model: {
                required: true
            }
        },
        mounted() {
            this.series[0] = {...this.series[0], ...{data: this.model.data}};
            this.chartOptions = {...this.chartOptions, ...{xaxis: {categories: this.model.name}}};
            this.chartOptions = {...this.chartOptions, ...{subtitle: {text: this.model.startDate+ " - " +this.model.endDate}}};
            this.chartOptions = {...this.chartOptions, ...{title: {text: this.model.title}}};
        },
        data: function () {
            return {
                series: [],
                chartOptions: {
                    plotOptions: {
                        bar: {
                            horizontal: true,
                            dataLabels: {
                                position: "top"
                            }
                        }
                    },
                     title: {
                      text: '',
                      margin: 10,
                        x: -20,
                        floating: false,
                        style: {
                            fontSize: "20px",
                        }
                    },
                    subtitle: {
                        text: "",
                        align: "left",
                        margin: 10,
                        x: -20,
                        floating: false,
                        style: {
                            fontSize: "15px",
                        }
                    },
                    xaxis: {
                        categories: []
                    }
                }
            };
        },
        methods: {
            updateTheme(e) {
                this.chartOptions = {
                    theme: {
                        palette: e.target.value
                    }
                };
            }
        }
    };
</script>
