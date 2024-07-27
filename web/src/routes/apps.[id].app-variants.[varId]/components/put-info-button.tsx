import useApi from '@/common/use-api';
import useAppVariant from '../hooks/use-app-variant';
import { useSnackbar } from 'notistack';
import useDialog from '@/common/use-dialog';
import { useFormik } from 'formik';
import * as Yup from 'yup';
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Stack,
  TextField,
} from '@mui/material';
import useFieldProperties from '@/routes/hooks/use-field-properties';
import { useEffect } from 'react';

export default function PutInfoButton() {
  const { data, mutate } = useAppVariant();
  const appVar = data?.data;

  const api = useApi();
  const { enqueueSnackbar } = useSnackbar();
  const { isOpen, onClose, onOpen } = useDialog();

  const formik = useFormik({
    initialValues: {
      title: appVar?.title ?? '',
      description: appVar?.description ?? '',
    },
    validationSchema: Yup.object({
      title: Yup.string().required('泳道名称不能为空'),
      description: Yup.string(),
    }),
    onSubmit: async values => {
      if (!appVar) throw new Error('cannot find app variant');

      await api.PUT('/api/app-variant/{id}', {
        params: {
          path: {
            id: appVar.id,
          },
        },
        body: {
          ...appVar,
          ...values,
        },
      });

      onClose();
      mutate();
      enqueueSnackbar('修改基本信息成功', { variant: 'success' });
    },
  });

  const { getFieldProperties } = useFieldProperties(formik);

  useEffect(() => {
    formik.resetForm();
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
                label="泳道名称"
                variant="standard"
                {...getFieldProperties('title')}
              />
              <TextField
                label="简介"
                multiline
                minRows={3}
                maxRows={5}
                variant="standard"
                {...getFieldProperties('description')}
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
