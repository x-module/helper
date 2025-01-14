### 1. 支持的统计图表类型

1. 折线图（Line Chart）：
    *
        * 用于显示数据随时间或类别的变化趋势。
    *
        * 示例：chart.LineChart。
2. 柱状图（Bar Chart）：
    * 用于比较不同类别的数据。
    * 示例：chart.BarChart。
3. 堆叠柱状图（Stacked Bar Chart）：
    * 用于显示多个数据系列的堆叠效果。
    * 示例：chart.StackedBarChart。
4. 饼图（Pie Chart）：
    * 用于显示各部分占总体的比例。
    * 示例：chart.PieChart。
5. 散点图（Scatter Chart）：
    * 用于显示两个变量之间的关系。
    * 示例：chart.ScatterChart。
6. 面积图（Area Chart）：
    * 类似于折线图，但填充折线下的区域。
    * 示例：chart.AreaChart。
7. 堆叠面积图（Stacked Area Chart）：
    * 多个面积图堆叠在一起，显示总量的变化趋势。
    * 示例：chart.StackedAreaChart。
8. 组合图（Combined Chart）：
    * 将多种图表类型（如折线图和柱状图）组合在一起。
    * 示例：chart.CombinedChart。
9. 箱线图（Box Plot）：
    * 用于显示数据的分布和离群值。
    * 示例：chart.BoxPlot。
10. 雷达图（Radar Chart）：
    * 用于显示多变量数据的对比。
    * 示例：chart.RadarChart。
11. 气泡图（Bubble Chart）：
    * 类似于散点图，但点的大小表示第三个变量。
    * 示例：chart.BubbleChart。
12. 蜡烛图（Candlestick Chart）：
    * 用于显示金融数据（如股票价格）。
    * 示例：chart.CandlestickChart。
13. 仪表盘图（Gauge Chart）：
    * 用于显示单一指标的进度或状态。
    * 示例：chart.GaugeChart。
14. 热力图（Heatmap Chart）：
    * 用于显示二维数据的密度或强度。
    * 示例：chart.HeatmapChart。
15. 直方图（Histogram Chart）：
    * 用于显示数据的分布情况。
    * 示例：chart.HistogramChart。

### 2.主要功能

1. 高度可定制：
    * 支持自定义颜色、字体、标签、图例等。
    * 可以调整图表的大小、边距、背景等。
2. 输出格式：
    * 支持生成 PNG、JPEG、SVG 等格式的图像。
3. 交互性：
    * 支持生成带有交互功能的 SVG 图表（如悬停提示）。
4. 轻量级：
   * 纯 Go 实现，无需依赖外部库。

### 3.支持的图表类型总结

| 图表类型  | 描述              | * 示例类名                 |
|-------|-----------------|------------------------|                                   
| 折线图   | 显示数据的变化趋势       | chart.LineChart        |               |
| 柱状图   | 比较不同类别的数据       | chart.BarChart         |
| 堆叠柱状图 | 显示多个数据系列的堆叠效果   | chart.StackedBarChart  |
| 饼图    | 显示各部分占总体的比例     | chart.PieChart         |
| 散点图   | 显示两个变量之间的关系     | chart.ScatterChart     |
| 面积图   | 填充折线下的区域        | chart.AreaChart        |
| 堆叠面积图 | 多个面积图堆叠显示总量变化   | chart.StackedAreaChart |
| 组合图   | 组合多种图表类型        | chart.CombinedChart    |
| 箱线图   | 显示数据分布和离群值      | chart.BoxPlot          |
| 雷达图   | 显示多变量数据的对比      | chart.RadarChart       |
| 气泡图   | 散点图，点的大小表示第三个变量 | chart.BubbleChart      |
| 蜡烛图   | 显示金融数据（如股票价格）   | chart.CandlestickChart |
| 仪表盘图  | 显示单一指标的进度或状态    | chart.GaugeChart       |
| 热力图   | 显示二维数据的密度或强度    | chart.HeatmapChart     |
| 直方图   | 显示数据的分布情况       | chart.HistogramChart   |

### 4.总结

go-chart 支持多种常见的统计图表类型，适 * 用于大多数数据可视化需求。如果需要更复杂的图表或功能，可以结合其他库（如 plotly 或 echarts）使用。

如果你有具体的图表需求或问题，欢迎随时告诉我！
开启新对话
