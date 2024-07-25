import useAppList from './hooks/use-app-list';
import DeleteAppButton from './components/delete-app-button';
import {
  IconButton,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from '@mui/material';
import VisibilityIcon from '@mui/icons-material/Visibility';
import { Link } from '@modern-js/runtime/router';

export default function AppList() {
  const { data } = useAppList();
  const apps = data?.data;

  return (
    <>
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell>ID</TableCell>
              <TableCell>App 名称</TableCell>
              <TableCell>简介</TableCell>
              <TableCell>域名匹配</TableCell>
              <TableCell>操作</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {apps?.map(app => (
              <TableRow key={app.id}>
                <TableCell component="th" scope="row">
                  {app.id}
                </TableCell>
                <TableCell>{app.title}</TableCell>
                <TableCell>{app.description}</TableCell>
                <TableCell>{app.hostnames.join(',')}</TableCell>
                <TableCell>
                  <Link to={`/apps/${app.id}`}>
                    <IconButton>
                      <VisibilityIcon />
                    </IconButton>
                  </Link>
                  <DeleteAppButton app={app} />
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
}
