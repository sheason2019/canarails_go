import Grid from '@mui/material/Unstable_Grid2';
import useAppList from './hooks/use-app-list';
import AppItem from './components/app-item';

export default function AppList() {
  const { data } = useAppList();



  return (
    <Grid container spacing={2}>
      {data?.data?.map(app => (
        <Grid xs={12} md={4}>
          <AppItem app={app} />
        </Grid>
      ))}
    </Grid>
  );
}
