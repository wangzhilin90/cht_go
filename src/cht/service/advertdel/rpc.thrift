//文章管理---广告图片管理---删除广告图片
namespace php Article.AdvertDel
namespace go  advertdel

struct AdvertDelRequestStruct{
	1:i32 id,
	2:string chengHuiTongTraceLog
}

struct AdvertDelResponseStruct{
	1:i32 status, //状态码
	2:string msg,
	3:i32 fid  	  //删除后返回图片地址
}

service AdvertDelThriftService {
	AdvertDelResponseStruct delAdvert (1:AdvertDelRequestStruct requestObj)
}

//		  if(!$id){
//            return false;
//        }
//		$fid = $this->db->result($this->db->c_sql(array (
//			'id' => $id
//		), "fid", '#@_advert_manage'));
//
//        if ($fid) {//删除图片
//            $upload = O('upload');
//            $upload->delete(array('id' => $fid));
//        }

//SELECT `fid` FROM #@_advert_manage WHERE `id`='1'
