import { useState, useEffect } from 'react';
import { GetMetrics } from "../wailsjs/go/main/App";
import {
  XAxis,
  YAxis,
  Tooltip,
  ResponsiveContainer,
  AreaChart,
  Area
} from 'recharts';

interface HistoryEntry {
  time: string;
  cpu: number;
  memory: number;
  netIn: number;
  netOut: number;
}

const formatBytes = (bytes: number, decimals = 2) => {
  if (bytes === 0) return '0 Bytes';
  const k = 1024;
  const dm = decimals < 0 ? 0 : decimals;
  const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i];
};

function App() {
  const [metrics, setMetrics] = useState<any>(null);
  const [history, setHistory] = useState<HistoryEntry[]>([]);
  const [lastNet, setLastNet] = useState<{ in: number, out: number } | null>(null);

  useEffect(() => {
    const interval = setInterval(() => {
      GetMetrics().then((res) => {
        setMetrics(res);
        const timestamp = new Date().toLocaleTimeString();

        // Calculate network rates (sum of all interfaces)
        const netData = res?.network?.Data?.network?.Interfaces || [];
        const totalIn = netData.reduce((acc: number, curr: any) => acc + curr.BytesIn, 0);
        const totalOut = netData.reduce((acc: number, curr: any) => acc + curr.BytesOut, 0);

        let rateIn = 0;
        let rateOut = 0;

        if (lastNet) {
          rateIn = Math.max(0, totalIn - lastNet.in);
          rateOut = Math.max(0, totalOut - lastNet.out);
        }
        setLastNet({ in: totalIn, out: totalOut });

        setHistory(prev => {
          const newEntry = {
            time: timestamp,
            cpu: res?.cpu?.Data?.cpu?.TotalUsage || 0,
            memory: (res?.memory?.Data?.memory?.Used / 1024 / 1024 / 1024) || 0,
            netIn: rateIn,
            netOut: rateOut,
          };
          const updated = [...prev, newEntry];
          return updated.slice(-30);
        });
      });
    }, 1000);
    return () => clearInterval(interval);
  }, [lastNet]);

  const cpuData = metrics?.cpu?.Data?.cpu;
  const memData = metrics?.memory?.Data?.memory;
  const diskData = metrics?.disk?.Data?.disk;
  const netData = metrics?.network?.Data?.network;

  return (
    <div className="grid-dashboard">
      {/* CPU CARD */}
      <div className="glass-card" style={{ gridColumn: 'span 2' }}>
        <div className="metric-label">CPU Performance</div>
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-end', marginBottom: '1rem' }}>
          <div className="metric-value" style={{ color: 'var(--accent-primary)' }}>
            {cpuData?.TotalUsage?.toFixed(1)}%
          </div>
          <div style={{ textAlign: 'right', color: 'var(--text-muted)', fontSize: '0.9rem' }}>
            {cpuData?.Frequency} MHz | {cpuData?.PerCore?.length} Cores
          </div>
        </div>
        <div style={{ height: '180px', width: '100%' }}>
          <ResponsiveContainer width="100%" height="100%">
            <AreaChart data={history}>
              <defs>
                <linearGradient id="colorCpu" x1="0" y1="0" x2="0" y2="1">
                  <stop offset="5%" stopColor="var(--accent-primary)" stopOpacity={0.3} />
                  <stop offset="95%" stopColor="var(--accent-primary)" stopOpacity={0} />
                </linearGradient>
              </defs>
              <Tooltip
                contentStyle={{ background: 'var(--card-surface)', border: '1px solid rgba(255,255,255,0.1)', borderRadius: '8px', backdropFilter: 'blur(10px)' }}
                itemStyle={{ color: 'var(--text-primary)' }}
              />
              <Area type="monotone" dataKey="cpu" stroke="var(--accent-primary)" fillOpacity={1} fill="url(#colorCpu)" isAnimationActive={false} strokeWidth={2} />
            </AreaChart>
          </ResponsiveContainer>
        </div>
      </div>

      {/* MEMORY CARD */}
      <div className="glass-card">
        <div className="metric-label">Memory Usage</div>
        <div className="metric-value" style={{ color: 'var(--state-healthy)' }}>
          {((memData?.Used / 1024 / 1024 / 1024) || 0).toFixed(1)} GB
        </div>
        <div className="progress-bar-bg">
          <div
            className="progress-bar-fill"
            style={{
              width: `${Math.min((memData?.Used / (memData?.Used + 1e9)) * 100, 100)}%`, // Simplified % logic for now
              backgroundColor: 'var(--state-healthy)'
            }}
          />
        </div>
        <div style={{ marginTop: '1.5rem', display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '0.5rem' }}>
          <div>
            <div className="metric-label" style={{ fontSize: '0.7rem' }}>Wired</div>
            <div style={{ fontSize: '0.9rem' }}>{formatBytes(memData?.Wired)}</div>
          </div>
          <div>
            <div className="metric-label" style={{ fontSize: '0.7rem' }}>Compressed</div>
            <div style={{ fontSize: '0.9rem' }}>{formatBytes(memData?.Compressed)}</div>
          </div>
          <div>
            <div className="metric-label" style={{ fontSize: '0.7rem' }}>Pressure</div>
            <div style={{ fontSize: '0.9rem', color: memData?.Pressure > 2 ? 'var(--state-critical)' : 'var(--state-healthy)' }}>
              {memData?.Pressure?.toFixed(0)}%
            </div>
          </div>
          <div>
            <div className="metric-label" style={{ fontSize: '0.7rem' }}>Swap</div>
            <div style={{ fontSize: '0.9rem' }}>{formatBytes(memData?.SwapUsed)}</div>
          </div>
        </div>
      </div>

      {/* DISK CARD */}
      <div className="glass-card">
        <div className="metric-label">Storage</div>
        {diskData?.Volumes?.slice(0, 2).map((v: any, i: number) => (
          <div key={i} style={{ marginBottom: '1rem' }}>
            <div style={{ display: 'flex', justifyContent: 'space-between', fontSize: '0.8rem', marginBottom: '0.25rem' }}>
              <span>{v.MountPoint}</span>
              <span style={{ color: 'var(--text-muted)' }}>{((v.Used / v.Total) * 100).toFixed(0)}%</span>
            </div>
            <div className="progress-bar-bg">
              <div
                className="progress-bar-fill"
                style={{
                  width: `${(v.Used / v.Total) * 100}%`,
                  backgroundColor: (v.Used / v.Total) > 0.9 ? 'var(--state-critical)' : 'var(--accent-primary)'
                }}
              />
            </div>
          </div>
        ))}
        <div style={{ display: 'flex', justifyContent: 'space-between', marginTop: '1rem' }}>
          <div>
            <div className="metric-label" style={{ fontSize: '0.7rem' }}>Read</div>
            <div style={{ fontSize: '0.9rem' }}>{formatBytes(diskData?.ReadBytesRate)}/s</div>
          </div>
          <div>
            <div className="metric-label" style={{ fontSize: '0.7rem' }}>Write</div>
            <div style={{ fontSize: '0.9rem' }}>{formatBytes(diskData?.WriteBytesRate)}/s</div>
          </div>
        </div>
      </div>

      {/* NETWORK CARD */}
      <div className="glass-card" style={{ gridColumn: 'span 2' }}>
        <div className="metric-label">Network Activity</div>
        <div style={{ display: 'flex', gap: '2rem', marginBottom: '1rem' }}>
          <div>
            <div style={{ fontSize: '0.7rem', color: 'var(--text-muted)' }}>DOWNLINK</div>
            <div style={{ color: 'var(--state-healthy)', fontWeight: 'bold' }}>{formatBytes(history[history.length - 1]?.netIn || 0)}/s</div>
          </div>
          <div>
            <div style={{ fontSize: '0.7rem', color: 'var(--text-muted)' }}>UPLINK</div>
            <div style={{ color: 'var(--accent-primary)', fontWeight: 'bold' }}>{formatBytes(history[history.length - 1]?.netOut || 0)}/s</div>
          </div>
        </div>
        <div style={{ height: '100px', width: '100%' }}>
          <ResponsiveContainer width="100%" height="100%">
            <AreaChart data={history}>
              <Tooltip
                contentStyle={{ background: 'var(--card-surface)', border: '1px solid rgba(255,255,255,0.1)', borderRadius: '8px' }}
                itemStyle={{ color: 'var(--text-primary)' }}
              />
              <Area type="monotone" dataKey="netIn" stroke="var(--state-healthy)" fill="var(--state-healthy)" fillOpacity={0.1} isAnimationActive={false} />
              <Area type="monotone" dataKey="netOut" stroke="var(--accent-primary)" fill="var(--accent-primary)" fillOpacity={0.1} isAnimationActive={false} />
            </AreaChart>
          </ResponsiveContainer>
        </div>
      </div>

      {/* OPEN PORTS */}
      <div className="glass-card" style={{ gridRow: 'span 1', maxHeight: '250px', overflowY: 'auto' }}>
        <div className="metric-label">Open Ports</div>
        <div style={{ fontSize: '0.8rem' }}>
          {netData?.OpenPorts?.slice(0, 10).map((p: any, i: number) => (
            <div key={i} style={{ display: 'flex', justifyContent: 'space-between', padding: '0.25rem 0', borderBottom: '1px solid rgba(255,255,255,0.05)' }}>
              <span style={{ color: 'var(--accent-primary)' }}>:{p.Port}</span>
              <span style={{ color: 'var(--text-muted)', overflow: 'hidden', textOverflow: 'ellipsis', whiteSpace: 'nowrap', maxWidth: '100px' }}>{p.Process}</span>
              <span>{p.Protocol}</span>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default App;
