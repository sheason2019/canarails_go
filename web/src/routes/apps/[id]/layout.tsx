import {
  Link,
  Outlet,
  useLocation,
  useNavigate,
} from '@modern-js/runtime/router';
import { Box, Typography, Tabs, Tab } from '@mui/material';
import useApp from './hooks/use-app';
import { useMemo } from 'react';

export default function Layout() {
  const { data } = useApp();
  const app = data?.data;
  const navigate = useNavigate();

  const location = useLocation();
  const tabValue = useMemo(() => {
    const tabName = location.pathname.split('/')[3];
    switch (tabName) {
      case 'app-variants':
        return 1;
      case 'logs':
        return 2;
      default:
        return 0;
    }
  }, [location.pathname]);

  return (
    <div>
      <Box sx={{ mt: 2, mb: 1 }}>
        <Typography variant="h4">{app?.title}</Typography>
        <Typography variant="body2" color="GrayText">
          ID {app?.id}
        </Typography>
      </Box>
      <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
        <Tabs value={tabValue}>
          <Tab label="基本信息" onClick={() => navigate(`/apps/${app?.id}`)} />
          <Tab
            label="流量泳道"
            onClick={() => navigate(`/apps/${app?.id}/app-variants`)}
          />
          <Tab
            label="日志信息"
            onClick={() => navigate(`/apps/${app?.id}/logs`)}
          />
        </Tabs>
      </Box>
      <Outlet />
    </div>
  );
}
