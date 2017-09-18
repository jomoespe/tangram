FROM scratch

LABEL project="The Tangram" \
      description="An Edge-side Universal Rendering service" \
      author="jose.morenoesteban@zooplus.com" \
      scope="Zooplus Exploration Days" 

ADD build/tangram /tangram

ENTRYPOINT ["./tangram"]
