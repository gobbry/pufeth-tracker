import React, { CSSProperties, PureComponent } from "react";
import axios from "axios";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from "recharts";
import Image from "next/image";

export default class PufferChart extends PureComponent {
  state = {
    data: [],
    loading: true,
  };

  intervalId: NodeJS.Timeout | null = null;

  componentDidMount() {
    this.fetchData();

    this.intervalId = setInterval(() => {
      this.fetchData();
    }, 12_000);
  }

  componentWillUnmount() {
    if (this.intervalId) {
      clearInterval(this.intervalId);
    }
  }

  fetchData = async () => {
    try {
      const response = await axios.get("http://localhost:3001/rate-series");
      const transformedData = response.data.rates
        .map((item: any) => ({
          timestamp: new Date(item.BlockTimestamp).toLocaleTimeString(),
          rate: parseFloat(item.ConversionRate),
          assets: parseFloat(Number(item.TotalAsset).toFixed(2)),
          supply: parseFloat(Number(item.TotalSupply).toFixed(2)),
        }))
        .reverse();
      this.setState({
        data: transformedData,
        loading: false,
      });
    } catch (err) {
      console.error(err);
      this.setState({ loading: false });
    }
  };

  getDataRange = (key: string) => {
    const values = this.state.data.map((item) => item[key]);
    const min = Math.min(...values);
    const max = Math.max(...values);
    const padding = (max - min) * 0.3; // 5% padding
    return [min - padding, max + padding];
  };

  margin = {
    top: 10,
    right: 50,
    left: 80,
    bottom: 0,
  };

  tooltipStyle = {
    backgroundColor: "#000000",
    border: "1px solid #befe38",
    borderRadius: "4px",
    padding: "10px",
  };

  tooltipLabelStyle = {
    color: "#befe38",
    marginBottom: "5px",
  };

  tooltipItemStyle = {
    color: "#befe38",
    padding: "3px 0",
  };

  chartStyle: CSSProperties = {
    position: "relative",
    background: "#101a3f",
    borderRadius: "8px",
    padding: "20px",
    marginBottom: "20px",
    height: "280px",
  };

  chartTitleStyle = {
    color: "#befe38",
    fontSize: "1.28rem",
    fontWeight: "888",
    marginBottom: "1rem",
    paddingLeft: "0.5rem",
  };

  chartContainerStyle = {
    position: "relative",
    background: "#101a3f",
    borderRadius: "8px",
    padding: "20px",
    marginBottom: "20px",
    overflow: "hidden",
  };

  chartBackgroundStyle = {
    position: "absolute",
    bottom: "10px",
    right: "10px",
    width: "150px",
    height: "auto",
    opacity: 0.1,
    pointerEvents: "none",
  };

  watermarkStyle = {
    position: "absolute",
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
    fontSize: "24px",
    fontWeight: "bold",
    color: "#ffffff",
    opacity: 0.1,
    pointerEvents: "none",
    zIndex: 0,
  };

  render() {
    if (this.state.loading) return <div>Loading...</div>;

    const CustomTooltip = ({ active, payload, label }: any) => {
      if (active && payload && payload.length) {
        return (
          <div style={this.tooltipStyle}>
            <p style={this.tooltipLabelStyle}>{label}</p>
            {payload.map((pld: any) => (
              <p key={pld.name} style={this.tooltipItemStyle}>
                {`${pld.name}: ${pld.value}`}
              </p>
            ))}
          </div>
        );
      }
      return null;
    };

    return (
      <div style={{ width: "100%" }} className="bg-[#fefffe] p-6 rounded-xl">
        <div style={this.chartStyle}>
          <h3 style={this.chartTitleStyle}>PufETH:ETH Conversion Rate</h3>
          <Image
            src="/big-puffer.webp"
            alt=""
            className="absolute bottom-2 right-1 w-52 opacity-15 pointer-events-none"
          />
          <div className="absolute inset-0 flex items-center justify-center pointer-events-none">
            <span className="text-white text-2xl font-bold opacity-10">
              PUFFER FINANCE
            </span>
          </div>

          <ResponsiveContainer width="100%" height={200}>
            <LineChart
              width={1200}
              height={200}
              data={this.state.data}
              syncId="anyId"
              margin={this.margin}
            >
              <CartesianGrid strokeDasharray="3 3" opacity={0.2} />
              <XAxis dataKey="timestamp" stroke="#fefffe" />
              <YAxis domain={this.getDataRange("rate")} stroke="#fefffe" />
              <Tooltip content={<CustomTooltip />} />
              <Line
                type="monotone"
                dataKey="rate"
                name="Rate"
                stroke="#befe38"
                dot={false}
              />
            </LineChart>
          </ResponsiveContainer>
        </div>

        <div style={this.chartStyle}>
          <h3 style={this.chartTitleStyle}>Total ETH in Puffer Vault</h3>
          <Image
            src="/big-puffer.webp"
            alt=""
            className="absolute bottom-2 right-1 w-52 opacity-15 pointer-events-none"
          />
          <div className="absolute inset-0 flex items-center justify-center pointer-events-none">
            <span className="text-white text-2xl font-bold opacity-10">
              PUFFER FINANCE
            </span>
          </div>

          <ResponsiveContainer width="100%" height={200}>
            <LineChart
              width={1200}
              height={200}
              data={this.state.data}
              syncId="anyId"
              margin={this.margin}
            >
              <CartesianGrid strokeDasharray="3 3" opacity={0.2} />
              <XAxis dataKey="timestamp" stroke="#fefffe" />
              <YAxis domain={this.getDataRange("assets")} stroke="#fefffe" />
              <Tooltip content={<CustomTooltip />} />
              <Line
                type="monotone"
                dataKey="assets"
                name="Assets"
                stroke="#befe38"
                dot={false}
              />
            </LineChart>
          </ResponsiveContainer>
        </div>

        <div style={this.chartStyle}>
          <h3 style={this.chartTitleStyle}>Total PufETH nLRT Minted</h3>
          <Image
            src="/big-puffer.webp"
            alt=""
            className="absolute bottom-2 right-1 w-52 opacity-15 pointer-events-none"
          />
          <div className="absolute inset-0 flex items-center justify-center pointer-events-none">
            <span className="text-white text-2xl font-bold opacity-10">
              PUFFER FINANCE
            </span>
          </div>

          <ResponsiveContainer width="100%" height={200}>
            <LineChart
              width={1200}
              height={200}
              data={this.state.data}
              syncId="anyId"
              margin={this.margin}
            >
              <CartesianGrid strokeDasharray="3 3" opacity={0.2} />
              <XAxis dataKey="timestamp" stroke="#fefffe" />
              <YAxis domain={this.getDataRange("supply")} stroke="#fefffe" />
              <Tooltip content={<CustomTooltip />} />
              <Line
                type="monotone"
                dataKey="supply"
                name="Supply"
                stroke="#befe38"
                dot={false}
              />
            </LineChart>
          </ResponsiveContainer>
        </div>
      </div>
    );
  }
}
