import useDialog from '@/common/use-dialog';
import useFieldProperties from '@/routes/hooks/use-field-properties';
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Stack,
  TextField,
} from '@mui/material';
import { useFormik } from 'formik';
import { useEffect } from 'react';
import * as Yup from 'yup';
import useAppVariant from '../hooks/use-app-variant';
import useApi from '@/common/use-api';
import { useSnackbar } from 'notistack';

export default function AddHeaderMatchButton() {
  const { data, mutate } = useAppVariant();
  const { isOpen, onClose, onOpen } = useDialog();
  const api = useApi();
  const { enqueueSnackbar } = useSnackbar();

  const appVar = data?.data;

  const formik = useFormik({
    initialValues: {
      header: '',
      value: '',
    },
    validationSchema: Yup.object({
      header: Yup.string().required('header 不能为空'),
      value: Yup.string().required('value 不能为空'),
    }),
    onSubmit: async values => {
      if (!appVar) throw new Error('cannot find app variant');

      await api.PUT('/api/app-variant/{id}', {
        params: {
          path: { id: appVar.id },
        },
        body: {
          ...appVar,
          matches: [...appVar.matches, values],
        },
      });

      enqueueSnackbar('新增标头匹配成功', { variant: 'success' });
      onClose();
      mutate();
    },
  });

  const { getFieldProperties } = useFieldProperties(formik);

  useEffect(() => {
    formik.resetForm();
  }, [isOpen]);

  return (
    <>
      <Button variant="contained" onClick={onOpen}>
        新增
      </Button>
      <Dialog open={isOpen} onClose={onClose} fullWidth>
        <form onSubmit={formik.handleSubmit}>
          <DialogTitle>新增标头匹配</DialogTitle>
          <DialogContent>
            <Stack spacing={2}>
              <TextField
                label="Header"
                variant="standard"
                {...getFieldProperties('header')}
              />
              <TextField
                label="Value"
                variant="standard"
                {...getFieldProperties('value')}
              />
            </Stack>
          </DialogContent>
          <DialogActions>
            <Button onClick={onClose}>取消</Button>
            <Button variant="contained" type="submit">
              提交
            </Button>
          </DialogActions>
        </form>
      </Dialog>
    </>
  );
}
