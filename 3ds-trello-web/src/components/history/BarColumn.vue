<template>
    <div class="animated fadeIn">
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
        <apexchart height="372" type="bar" :options="chartOptions" :series="series"></apexchart>
    </div>
</template>

<script>
    export default {
        name: "BarTotal",
        props: {
            model: {
                required: true,
            },
        },
        mounted() {
            this.series[0] = {...this.series[0], ...{data: this.model.data}};
            this.chartOptions = {...this.chartOptions, ...{xaxis: {categories: this.model.name}}};
        },
        data: function () {
            return {
                chartOptions: {
                    chart: {
                        id: "vuechart-example"
                    },
                    xaxis: {
                        categories: []
                    },
                    plotOptions: {
                        bar: {
                            distributed: true,
                            dataLabels: {
                                position: "top"
                            }
                        }
                    },
                    subtitle: {
                        text: "Total",
                        align: "left",
                        margin: 10,
                        offsetX: 0,
                        offsetY: 0,
                        floating: false,
                        style: {
                            fontSize: "16px",
                            color: "#000000"
                        }
                    }
                },
                series: [
                    {
                        name: "Total",
                        data: []
                    }
                ]
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
