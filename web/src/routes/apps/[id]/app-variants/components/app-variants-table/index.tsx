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
import useAppVariants from './hooks/use-app-variants';
import { Link } from '@modern-js/runtime/router';
import DeleteAppVariantButton from './delete-app-variant-button';

export default function AppVariantsTable() {
  const { data } = useAppVariants();

  const dataSource = data?.data;

  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>ID</TableCell>
            <TableCell>名称</TableCell>
            <TableCell>简介</TableCell>
            <TableCell>操作</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {dataSource?.map(appVar => (
            <TableRow key={appVar.id}>
              <TableCell component="th" scope="row">
                {appVar.id}
              </TableCell>
              <TableCell>{appVar.title}</TableCell>
              <TableCell>{appVar.description}</TableCell>
              <TableCell>
                <Link to={`/apps/${appVar.appId}/app-variants/${appVar.id}`}>
                  <IconButton>
                    <VisibilityIcon />
                  </IconButton>
                </Link>
                <DeleteAppVariantButton appVar={appVar} />
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
