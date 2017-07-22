package objects

type BVH_Node struct {
    left, right *Hitable
    box AABB
}

func NewBVH_Node() BVH_Node {

}


func (b *BVH_Node) Hit(r Ray, tMin float64, tMax float64)(bool, Hit) {
    var rec Hit
    if t, _ := b.box.Hit(r, tMin, tMax); t {
        left_rec, hit_left := b.left.Hit(r, tMin, tMax)
        right_rec, hit_right := b.right.Hit(r, tMin, tMax)
        if hit_left && hit_right {
            if left_rec.T < right_rec.T
              rec = left_rec
            else
              rec = right_rec
            return true, rec
        }
        else if hit_left {
            rec = left_rec
            return true, rec
        }
        else if hit_right {
            rec = right_rec
            return true, rec
        }
        else
          return false, Hit{}
    }
    else
      return false, Hit{}
}

func (b* BVH_Node) BoundingBox(t0, t1 float)(bool, AABB) {
    return true, b.box
}
