namespace go advertmanagethriftservice

struct AdvertManageRequestStruct {
    1:i32 type,
    2:i32 limit,
    3:string chengHuiTongTraceLog
}

struct AdvertManageStruct {
    1:i32 id,
    2:i32 type, //���λ�ã���Ӧ�ֵ�45��ֵ����ԭֵ��1ע�᣻2��¼��3��ҳ���Ͻǣ�4��ҳ�·���5��Ŀ�б�6��Ŀ���飻7�������ģ�
    3:string img,
    4:string adverturl,
    5:string title,
    6:i32 addtime,
    7:i32 adduser,
    8:i32 status,
    9:i32 fid,
    10:i32 starttime,
    11:i32 endtime,
}

struct AdvertManageResponseStruct {
	1:i32 status
	2:string msg
    3:list<AdvertManageStruct> advertManageList
}

service AdvertManageThriftService {
    AdvertManageResponseStruct getAdvertManage(1:AdvertManageRequestStruct requestObj),
    //��ѯjl_banner_manage��
    //where������starttime <=  time() and endtime >= time() 
	//type = type
    //order by����addtime DESC
    //���limit=1ȡһ�����ݣ���������ȡ��������
}