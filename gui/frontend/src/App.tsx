import { useState, useEffect } from 'react';
import { GetMetrics } from "../wailsjs/go/main/App";
import { 
    Container, 
    Typography, 
    Grid, 
    Card, 
    CardContent, 
    Box, 
    ThemeProvider, 
    createTheme, 
    CssBaseline,
    LinearProgress,
    Divider
} from '@mui/material';
import { 
    XAxis, 
    YAxis, 
    CartesianGrid, 
    Tooltip, 
    ResponsiveContainer,
    AreaChart,
    Area
} from 'recharts';

const darkTheme = createTheme({
    palette: {
        mode: 'dark',
        primary: {
            main: '#7d56f4',
        },
        background: {
            default: '#1b2636',
            paper: '#253347',
        },
    },
});

function App() {
    const [metrics, setMetrics] = useState<any>(null);
    const [history, setHistory] = useState<any[]>([]);

    useEffect(() => {
        const interval = setInterval(() => {
            GetMetrics().then((newMetrics) => {
                setMetrics(newMetrics);
                const timestamp = new Date().toLocaleTimeString();
                
                setHistory(prev => {
                    const newEntry = {
                        time: timestamp,
                        cpu: newMetrics?.cpu?.Data?.cpu?.TotalUsage || 0,
                        memory: (newMetrics?.memory?.Data?.memory?.Used / 1024 / 1024 / 1024) || 0,
                    };
                    const updated = [...prev, newEntry];
                    return updated.slice(-30); // Keep last 30 seconds
                });
            });
        }, 1000);
        return () => clearInterval(interval);
    }, []);

    const cpuData = metrics?.cpu?.Data?.cpu;
    const memData = metrics?.memory?.Data?.memory;

    return (
        <ThemeProvider theme={darkTheme}>
            <CssBaseline />
            <Container maxWidth="lg" sx={{ py: 4 }}>
                <Typography variant="h4" component="h1" gutterBottom sx={{ color: 'primary.main', fontWeight: 'bold' }}>
                    mac-monitor Dashboard
                </Typography>

                <Grid container spacing={3}>
                    {/* CPU Chart */}
                    <Grid size={{ xs: 12, md: 8 }}>
                        <Card sx={{ height: '400px' }}>
                            <CardContent sx={{ height: '100%' }}>
                                <Typography variant="h6" gutterBottom>CPU Usage History (%)</Typography>
                                <ResponsiveContainer width="100%" height="90%">
                                    <AreaChart data={history}>
                                        <defs>
                                            <linearGradient id="colorCpu" x1="0" y1="0" x2="0" y2="1">
                                                <stop offset="5%" stopColor="#7d56f4" stopOpacity={0.8}/>
                                                <stop offset="95%" stopColor="#7d56f4" stopOpacity={0}/>
                                            </linearGradient>
                                        </defs>
                                        <CartesianGrid strokeDasharray="3 3" stroke="#3d4f6a" />
                                        <XAxis dataKey="time" stroke="#ccc" fontSize={10} />
                                        <YAxis domain={[0, 100]} stroke="#ccc" fontSize={10} />
                                        <Tooltip contentStyle={{ backgroundColor: '#253347', border: 'none' }} />
                                        <Area type="monotone" dataKey="cpu" stroke="#7d56f4" fillOpacity={1} fill="url(#colorCpu)" isAnimationActive={false} />
                                    </AreaChart>
                                </ResponsiveContainer>
                            </CardContent>
                        </Card>
                    </Grid>

                    {/* CPU Real-time Stats */}
                    <Grid size={{ xs: 12, md: 4 }}>
                        <Card sx={{ height: '400px', display: 'flex', flexDirection: 'column' }}>
                            <CardContent>
                                <Typography variant="h6" gutterBottom>Real-time CPU</Typography>
                                <Box sx={{ mb: 3 }}>
                                    <Typography variant="body2" color="text.secondary">Total Usage</Typography>
                                    <Typography variant="h4">{cpuData?.TotalUsage?.toFixed(1)}%</Typography>
                                    <LinearProgress variant="determinate" value={cpuData?.TotalUsage || 0} sx={{ height: 10, borderRadius: 5, mt: 1 }} />
                                </Box>
                                <Divider sx={{ my: 2 }} />
                                <Typography variant="body2" color="text.secondary">Frequency</Typography>
                                <Typography variant="h5">{cpuData?.Frequency} MHz</Typography>
                                <Typography variant="body2" color="text.secondary" sx={{ mt: 2 }}>Cores: {cpuData?.PerCore?.length}</Typography>
                            </CardContent>
                            <Box sx={{ p: 2, flexGrow: 1, overflowY: 'auto' }}>
                                <Grid container spacing={1}>
                                    {cpuData?.PerCore?.map((usage: number, i: number) => (
                                        <Grid size={4} key={i}>
                                            <Box sx={{ bgcolor: 'background.default', p: 1, borderRadius: 1, textAlign: 'center' }}>
                                                <Typography variant="caption" display="block">Core {i}</Typography>
                                                <Typography variant="body2" fontWeight="bold">{usage.toFixed(0)}%</Typography>
                                            </Box>
                                        </Grid>
                                    ))}
                                </Grid>
                            </Box>
                        </Card>
                    </Grid>

                    {/* Memory Chart */}
                    <Grid size={{ xs: 12, md: 8 }}>
                        <Card sx={{ height: '400px' }}>
                            <CardContent sx={{ height: '100%' }}>
                                <Typography variant="h6" gutterBottom>Memory Usage (GB)</Typography>
                                <ResponsiveContainer width="100%" height="90%">
                                    <AreaChart data={history}>
                                        <defs>
                                            <linearGradient id="colorMem" x1="0" y1="0" x2="0" y2="1">
                                                <stop offset="5%" stopColor="#faacac" stopOpacity={0.8}/>
                                                <stop offset="95%" stopColor="#faacac" stopOpacity={0}/>
                                            </linearGradient>
                                        </defs>
                                        <CartesianGrid strokeDasharray="3 3" stroke="#3d4f6a" />
                                        <XAxis dataKey="time" stroke="#ccc" fontSize={10} />
                                        <YAxis stroke="#ccc" fontSize={10} />
                                        <Tooltip contentStyle={{ backgroundColor: '#253347', border: 'none' }} />
                                        <Area type="monotone" dataKey="memory" stroke="#faacac" fillOpacity={1} fill="url(#colorMem)" isAnimationActive={false} />
                                    </AreaChart>
                                </ResponsiveContainer>
                            </CardContent>
                        </Card>
                    </Grid>

                    {/* Memory Real-time Stats */}
                    <Grid size={{ xs: 12, md: 4 }}>
                        <Card sx={{ height: '400px' }}>
                            <CardContent>
                                <Typography variant="h6" gutterBottom>Real-time Memory</Typography>
                                <Box sx={{ mb: 3 }}>
                                    <Typography variant="body2" color="text.secondary">Used</Typography>
                                    <Typography variant="h4">{(memData?.Used / 1024 / 1024 / 1024)?.toFixed(2)} GB</Typography>
                                </Box>
                                <Box sx={{ mb: 3 }}>
                                    <Typography variant="body2" color="text.secondary">Wired</Typography>
                                    <Typography variant="h5">{(memData?.Wired / 1024 / 1024 / 1024)?.toFixed(2)} GB</Typography>
                                </Box>
                                <Box sx={{ mb: 3 }}>
                                    <Typography variant="body2" color="text.secondary">Compressed</Typography>
                                    <Typography variant="h5">{(memData?.Compressed / 1024 / 1024 / 1024)?.toFixed(2)} GB</Typography>
                                </Box>
                                <Divider sx={{ my: 2 }} />
                                <Typography variant="body2" color="text.secondary">Memory Pressure</Typography>
                                <Typography variant="h4">{memData?.Pressure?.toFixed(0)}</Typography>
                                <LinearProgress 
                                    variant="determinate" 
                                    value={Math.min(memData?.Pressure * 25, 100) || 0}
                                    color={memData?.Pressure > 2 ? "error" : "success"}
                                    sx={{ height: 10, borderRadius: 5, mt: 1 }} 
                                />
                            </CardContent>
                        </Card>
                    </Grid>
                </Grid>
            </Container>
        </ThemeProvider>
    );
}

export default App;
