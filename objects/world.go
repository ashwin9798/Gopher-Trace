package objects

type World struct {
    Elements []Hitable
}

func (w *World) Add(h Hitable) {
	w.Elements = append(w.Elements, h)
}

func (w *World) AddAll(hitables ...Hitable) {
	for _, h := range hitables {
		w.Add(h)
	}
}

func (w *World) BoundingBox(t0, t1 float64)(bool, AABB) {
    if len(w.Elements) < 1
      return false
    tempBox := AABB{}
    var box AABB
    first_true, _ := w.Elements[0].BoundingBox(t0, t1, tempBox)

    if !first_true
      return false
    else
      box = tempBox

    for i:=1; i < len(w.Elements); i++ {
        true, _ := w.Elements[0].BoundingBox(t0,t1,tempBox))
        if(true){
          box = surroundingBox(box, tempBox)
        }
        else
          return false, AABB{}
    }
    return true, box
}

func (w *World) Hit(r Ray, tMin float64, tMax float64) (bool, Hit) {
    hitAnything := false
    closestSoFar := tMax
    record := Hit{}

    for _, element := range w.Elements {
        hit, tempRecord := element.Hit(r, tMin, closestSoFar)
        if hit {
            hitAnything = true
            closestSoFar = tempRecord.T
            record = tempRecord
        }
    }
    return hitAnything, record
}
