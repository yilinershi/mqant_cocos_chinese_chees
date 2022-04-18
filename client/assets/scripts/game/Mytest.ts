
import { _decorator, Component, Node } from 'cc';
const { ccclass, property } = _decorator;

 
@ccclass('Mytest')
export class Mytest extends Component {
    start () {
        let sex=pb.Sex.Male
        console.log(sex.toString())
    }
}

