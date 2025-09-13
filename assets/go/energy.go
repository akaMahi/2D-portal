components {
  id: "energy"
  component: "/assets/scripts/energy.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"energy\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlas/props.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.2
  }
  scale {
    x: 0.5
    y: 0.5
  }
}
embedded_components {
  id: "co"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"lava\"\n"
  "mask: \"default\"\n"
  "mask: \"player\"\n"
  "mask: \"box\"\n"
  "mask: \"feet\"\n"
  "mask: \"head\"\n"
  "mask: \"proj\"\n"
  "mask: \"white_wall\"\n"
  "mask: \"black_wall\"\n"
  "mask: \"op\"\n"
  "mask: \"bp\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "    id: \"circle\"\n"
  "  }\n"
  "  data: 7.5\n"
  "}\n"
  ""
}
embedded_components {
  id: "s"
  type: "sound"
  data: "sound: \"/assets/SFX/energy_ball.wav\"\n"
  "gain: 0.25\n"
  "speed: 2.0\n"
  ""
}
