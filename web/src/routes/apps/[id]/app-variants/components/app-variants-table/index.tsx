import {
  Chip,
  IconButton,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography,
} from '@mui/material';
import VisibilityIcon from '@mui/icons-material/Visibility';
import useAppVariants from './hooks/use-app-variants';
import DeleteAppVariantButton from './components/delete-app-variant-button';
import StyledLink from '@/common/styled-link';
import ControlMenuButton from './components/control-menu-button';
import useApp from '../../../hooks/use-app';

export default function AppVariantsTable() {
  const { data: appData } = useApp();
  const app = appData?.data;

  const { data: appVarData } = useAppVariants();
  const appVariants = appVarData?.data;

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
          {appVariants?.map(appVar => (
            <TableRow key={appVar.id}>
              <TableCell component="th" scope="row">
                {appVar.id}
                {appVar.id === app?.defaultVariantId && (
                  <Chip
                    sx={{ ml: 1 }}
                    label={<Typography variant="caption">默认泳道</Typography>}
                    size="small"
                    color="info"
                  />
                )}
              </TableCell>
              <TableCell>{appVar.title}</TableCell>
              <TableCell>{appVar.description}</TableCell>
              <TableCell>
                <StyledLink
                  to={`/apps/${appVar.appId}/app-variants/${appVar.id}`}
                >
                  <IconButton>
                    <VisibilityIcon />
                  </IconButton>
                </StyledLink>
                <DeleteAppVariantButton appVar={appVar} />
                <ControlMenuButton appVar={appVar} />
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
