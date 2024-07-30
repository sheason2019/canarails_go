import useApi from '@/common/use-api';
import useDialog from '@/common/use-dialog';
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
import { useSnackbar } from 'notistack';
import * as Yup from 'yup';
import useApp from '../hooks/use-app';
import { useEffect } from 'react';
import useFieldProperties from '@/routes/hooks/use-field-properties';

export default function PutAppInfoButton() {
  const { data, mutate } = useApp();
  const app = data?.data;

  const api = useApi();
  const { enqueueSnackbar } = useSnackbar();
  const { isOpen, onClose, onOpen } = useDialog();
  const formik = useFormik({
    initialValues: {
      title: app?.title ?? '',
      description: app?.description ?? '',
    },
    validationSchema: Yup.object({
      title: Yup.string().required('App 名称不能为空'),
      description: Yup.string(),
    }),
    onSubmit: async values => {
      if (!app) throw new Error('cannot find app');

      await api.PUT('/api/app/{id}', {
        params: {
          path: { id: Number(app?.id) },
        },
        body: {
          ...app,
          ...values,
        },
      });

      enqueueSnackbar('修改基本信息成功', { variant: 'success' });
      mutate();
      onClose();
    },
  });

  const { getFieldProperties } = useFieldProperties(formik);

  useEffect(() => {
    formik.resetForm({
      values: {
        title: app?.title ?? '',
        description: app?.description ?? '',
      },
    });
  }, [isOpen]);

  return (
    <>
      <Button variant="contained" onClick={onOpen}>
        修改
      </Button>
      <Dialog open={isOpen} onClose={onClose} fullWidth>
        <form onSubmit={formik.handleSubmit}>
          <DialogTitle>修改基本信息</DialogTitle>
          <DialogContent>
            <Stack spacing={2}>
              <TextField
                label="App名称"
                variant="standard"
                {...getFieldProperties('title')}
              />
              <TextField
                label="App简介"
                variant="standard"
                multiline
                minRows={3}
                maxRows={5}
                {...getFieldProperties('description')}
              />
            </Stack>
          </DialogContent>
          <DialogActions>
            <Button>取消</Button>
            <Button variant="contained" type="submit">
              提交
            </Button>
          </DialogActions>
        </form>
      </Dialog>
    </>
  );
}
